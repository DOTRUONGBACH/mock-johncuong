package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.30

import (
	"context"
	"mock-project/graphql/graph"
	"mock-project/graphql/graph/model"
	"mock-project/helper"
	"mock-project/middleware"
	pb "mock-project/pb/proto"
	"net/http"
	"strconv"

	"google.golang.org/grpc/metadata"
	"google.golang.org/protobuf/types/known/timestamppb"
)

// CreateBooking is the resolver for the createBooking field.
func (r *bookingOpsResolver) CreateBooking(ctx context.Context, obj *model.BookingOps, input *model.CreateBookingInput) (*model.CustomBookingResponse, error) {
	user := middleware.GetUserFromContext(ctx)

	// get flight to check booking time is valid
	fl, err := r.flightClient.GetFlight(ctx, &pb.GetFlightRequest{Name: input.FlightName})
	if err != nil {
		return nil, helper.GqlResponse(err.Error(), http.StatusInternalServerError)
	}

	if !helper.CheckCreateBooking(fl.DepartureDate.AsTime()) {
		return nil, helper.GqlResponse("you can't make booking in 4 hours before departure time", http.StatusForbidden)
	}

	// Guest Booking
	if user == nil {
		// send metadata to check if user is logged in
		md := metadata.Pairs("user", "0")
		ctx = metadata.NewOutgoingContext(ctx, md)

		err := helper.CheckBookingInput(input)
		if err != nil {
			return nil, helper.GqlResponse(err.Error(), http.StatusBadRequest)
		}

		// call gRPC booking
		conVTime, _ := helper.ConvertTime(helper.CheckNil(input.DateOfBirth))
		res, err := r.bookingClient.CreateBooking(ctx, &pb.CreateBookingRequest{
			FlightName:     input.FlightName,
			Email:          helper.CheckNil(input.Email),
			PhoneNumber:    helper.CheckNil(input.PhoneNumber),
			Name:           helper.CheckNil(input.Name),
			Address:        helper.CheckNil(input.Address),
			DateOfBirth:    timestamppb.New(conVTime),
			IdentifyNumber: helper.CheckNil(input.IdentifyNumber),
			MemberCode:     helper.CheckNil(input.MemberCode),
			TicketType:     int64(input.TicketType),
		})
		if err != nil {
			return nil, helper.GqlResponse(err.Error(), http.StatusInternalServerError)
		}

		// graphQL response
		resQ := &model.CustomCreateBookingResponse{
			BookingCode: &res.Code,
			BookingDate: helper.GetTimePointer(res.CreatedAt.AsTime()),
			CustomerID:  helper.GetIntPointer(int(res.CustomerId)),
			FlightID:    helper.GetIntPointer(int(res.FlightId)),
			Status:      &res.Status,
		}

		resQB := &model.CustomBookingResponse{
			Code: http.StatusCreated,
			Data: resQ,
		}
		return resQB, nil
	}

	// User booking
	if helper.CheckAdmin(user) == false {
		// send metadata to gRPC server can receive customerId of user
		md := metadata.Pairs("user", strconv.Itoa(user.CustomerID))
		ctx = metadata.NewOutgoingContext(ctx, md)
		// if user, only flight name and ticket type needed
		res, err := r.bookingClient.CreateBooking(ctx, &pb.CreateBookingRequest{
			FlightName: input.FlightName,
			TicketType: int64(input.TicketType),
		})
		if err != nil {
			return nil, helper.GqlResponse(err.Error(), http.StatusInternalServerError)
		}

		// graphQL response
		resQ := &model.CustomCreateBookingResponse{
			BookingCode: &res.Code,
			BookingDate: helper.GetTimePointer(res.CreatedAt.AsTime()),
			CustomerID:  helper.GetIntPointer(int(res.CustomerId)),
			FlightID:    helper.GetIntPointer(int(res.FlightId)),
			Status:      &res.Status,
		}

		resQB := &model.CustomBookingResponse{
			Code: http.StatusCreated,
			Data: resQ,
		}
		return resQB, nil
	}

	// Admin make booking :)
	return nil, helper.GqlResponse("admin cannot make booking", http.StatusForbidden)
}

// CancelBooking is the resolver for the cancelBooking field.
func (r *bookingOpsResolver) CancelBooking(ctx context.Context, obj *model.BookingOps, input *model.CancelBookingInput) (*model.CustomCancelResponse, error) {
	user := middleware.GetUserFromContext(ctx)
	if user == nil {
		return nil, helper.GqlResponse("please contact Airline (1800-000-000) to cancel booking", http.StatusForbidden)
	}

	if helper.CheckAdmin(user) == false {
		// check if booking code belongs to user
		booking, _ := r.bookingClient.GetBooking(ctx, &pb.GetBookingRequest{Code: input.BookingCode})
		if booking.CustomerId != int64(user.CustomerID) {
			return nil, helper.GqlResponse("this code wasn't booked by this account", http.StatusForbidden)
		}

		fl, err := r.flightClient.GetFlightById(ctx, &pb.GetFlightRequest{Id: booking.FlightId})
		if err != nil {
			return nil, helper.GqlResponse(err.Error(), http.StatusInternalServerError)
		}

		if !helper.CheckCancelBooking(fl.DepartureDate.AsTime()) {
			return nil, helper.GqlResponse("you can't cancel booking in 48 hours before departure time", http.StatusForbidden)
		}

		// call booking gRPC
		res, err := r.bookingClient.CancelBooking(ctx, &pb.CancelBookingRequest{BookingCode: input.BookingCode})
		if err != nil {
			return nil, helper.GqlResponse(err.Error(), http.StatusInternalServerError)
		}

		// graphQL response
		data := &model.CancelResponse{
			BookingCode: &res.Code,
			CancelDate:  helper.GetTimePointer(res.UpdatedAt.AsTime()),
			CustomerID:  helper.GetIntPointer(int(res.CustomerId)),
			FlightID:    helper.GetIntPointer(int(res.FlightId)),
			Status:      &res.Status,
		}

		resQ := &model.CustomCancelResponse{
			Code: http.StatusNoContent,
			Data: data,
		}
		return resQ, nil
	}
	return nil, helper.GqlResponse("admin cannot cancel booking", http.StatusForbidden)
}

// ViewBooking is the resolver for the viewBooking field.
func (r *bookingQueryResolver) ViewBooking(ctx context.Context, obj *model.BookingQuery, bookingCode string, identifyNumber string) (*model.CustomViewBookingResponse, error) {
	if identifyNumber == "" {
		return nil, helper.GqlResponse("identify number is required!", http.StatusBadRequest)
	}

	booking, err := r.bookingClient.GetBooking(ctx, &pb.GetBookingRequest{Code: bookingCode})
	if err != nil {
		return nil, helper.GqlResponse(err.Error(), http.StatusInternalServerError)
	}

	customer, err := r.customerClient.GetCustomer(ctx, &pb.GetCustomerRequest{Id: booking.CustomerId})
	if err != nil {
		return nil, helper.GqlResponse(err.Error(), http.StatusInternalServerError)
	}

	if customer.IdentifyNumber != identifyNumber {
		return nil, helper.GqlResponse("this booking code wasn't booked by provided identify number", http.StatusBadRequest)
	}

	flight, err := r.flightClient.GetFlightById(ctx, &pb.GetFlightRequest{Id: booking.FlightId})
	if err != nil {
		return nil, helper.GqlResponse(err.Error(), http.StatusInternalServerError)
	}

	customerQ := &model.Customer{
		ID:          int(customer.Id),
		Name:        customer.Name,
		Email:       customer.Email,
		Address:     customer.Address,
		PhoneNumber: customer.PhoneNumber,
		MemberCode:  customer.MemberCode,
	}

	flightQ := &model.Flight{
		ID:            int(flight.Id),
		Name:          flight.Name,
		From:          flight.From,
		To:            flight.To,
		DepartureDate: flight.DepartureDate.AsTime(),
		ArrivalDate:   flight.ArrivalDate.AsTime(),
		Status:        flight.Status,
	}

	data := &model.ViewBookingResponse{
		BookingCode:   &bookingCode,
		BookingDate:   helper.GetTimePointer(booking.UpdatedAt.AsTime()),
		BookingStatus: &booking.Status,
		Customer:      customerQ,
		Flight:        flightQ,
	}
	resQ := &model.CustomViewBookingResponse{
		Code: http.StatusOK,
		Data: data,
	}

	return resQ, nil
}

// BookingHistory is the resolver for the bookingHistory field.
func (r *bookingQueryResolver) BookingHistory(ctx context.Context, obj *model.BookingQuery, page *int, limit *int) (*model.CustomBookingHistoryResponse, error) {
	user := middleware.GetUserFromContext(ctx)

	if user == nil {
		return nil, helper.GqlResponse("guest cannot see booking history", http.StatusForbidden)
	}

	if helper.CheckAdmin(user) == false {
		req := &pb.GetBookingHistoryRequest{
			CustomerId: int64(user.CustomerID),
			Page:       int64(*page),
			Limit:      int64(*limit),
		}

		history, err := r.bookingClient.GetBookingHistory(ctx, req)
		if err != nil {
			return nil, helper.GqlResponse(err.Error(), http.StatusInternalServerError)
		}

		var bookingHRes []*model.ViewBookingResponse
		for _, bookingHistory := range history.BookingList {
			data := &model.ViewBookingResponse{}
			fl, err := r.flightClient.GetFlightById(ctx, &pb.GetFlightRequest{Id: bookingHistory.FlightId})
			if err != nil {
				return nil, helper.GqlResponse(err.Error(), http.StatusInternalServerError)
			}

			flightQ := &model.Flight{
				ID:            int(fl.Id),
				Name:          fl.Name,
				From:          fl.From,
				To:            fl.To,
				DepartureDate: fl.DepartureDate.AsTime(),
				ArrivalDate:   fl.ArrivalDate.AsTime(),
				Status:        fl.Status,
			}

			customer, err := r.customerClient.GetCustomer(ctx, &pb.GetCustomerRequest{Id: bookingHistory.CustomerId})
			if err != nil {
				return nil, helper.GqlResponse(err.Error(), http.StatusInternalServerError)
			}

			customerQ := &model.Customer{
				ID:          int(customer.Id),
				Name:        customer.Name,
				Email:       customer.Email,
				Address:     customer.Address,
				PhoneNumber: customer.PhoneNumber,
				MemberCode:  customer.MemberCode,
			}

			data.BookingCode = &bookingHistory.Code
			data.BookingDate = helper.GetTimePointer(bookingHistory.UpdatedAt.AsTime())
			data.BookingStatus = &bookingHistory.Status
			data.Flight = flightQ
			data.Customer = customerQ

			bookingHRes = append(bookingHRes, data)
		}

		resQ := &model.CustomBookingHistoryResponse{
			Code:  http.StatusOK,
			Data:  bookingHRes,
			Page:  helper.GetIntPointer(int(history.Page)),
			Limit: limit,
			Total: helper.GetIntPointer(int(history.Total)),
		}
		return resQ, nil
	}
	return nil, nil
}

// Bookings is the resolver for the bookings field.
func (r *bookingQueryResolver) Bookings(ctx context.Context, obj *model.BookingQuery, page *int, limit *int) (*model.BookingResponse, error) {
	user := middleware.GetUserFromContext(ctx)
	if user == nil {
		return nil, helper.GqlResponse("guest cannot see booking history", http.StatusForbidden)
	}

	if !helper.CheckAdmin(user) {
		return nil, helper.GqlResponse("only admin can see all bookings", http.StatusForbidden)
	}

	req := &pb.ListBookingRequest{
		Page:  int64(*page),
		Limit: int64(*limit),
	}

	// grpc call
	bookings, err := r.bookingClient.ListBooking(ctx, req)
	if err != nil {
		return nil, helper.GqlResponse(err.Error(), http.StatusInternalServerError)
	}

	// append data to response
	var bookingRes []*model.BookingFullResponse
	for _, booking := range bookings.BookingList {
		fl, err := r.flightClient.GetFlightById(ctx, &pb.GetFlightRequest{Id: booking.FlightId})
		if err != nil {
			return nil, helper.GqlResponse(err.Error(), http.StatusInternalServerError)
		}

		flightQ := &model.Flight{
			ID:                   int(fl.Id),
			Name:                 fl.Name,
			From:                 fl.From,
			To:                   fl.To,
			DepartureDate:        fl.DepartureDate.AsTime(),
			ArrivalDate:          fl.ArrivalDate.AsTime(),
			AvailableFirstSlot:   int(fl.AvailableFirstSlot),
			AvailableEconomySlot: int(fl.AvailableEconomySlot),
			Status:               fl.Status,
			CreatedAt:            helper.GetTimePointer(fl.CreatedAt.AsTime()),
			UpdatedAt:            helper.GetTimePointer(fl.UpdatedAt.AsTime()),
		}

		customer, err := r.customerClient.GetCustomer(ctx, &pb.GetCustomerRequest{Id: booking.CustomerId})
		if err != nil {
			return nil, helper.GqlResponse(err.Error(), http.StatusInternalServerError)
		}

		customerQ := &model.Customer{
			ID:             int(customer.Id),
			Name:           customer.Name,
			Email:          customer.Email,
			Address:        customer.Address,
			PhoneNumber:    customer.PhoneNumber,
			IdentifyNumber: customer.IdentifyNumber,
			DateOfBirth:    customer.DateOfBirth.AsTime().Format("2006-01-02"),
			MemberCode:     customer.MemberCode,
			CreatedAt:      helper.GetTimePointer(customer.CreatedAt.AsTime()),
			UpdatedAt:      helper.GetTimePointer(customer.UpdatedAt.AsTime()),
		}

		bookingFullRes := &model.BookingFullResponse{
			ID:        int(booking.Id),
			Customer:  customerQ,
			Flight:    flightQ,
			Code:      booking.Code,
			Status:    booking.Status,
			TicketID:  int(booking.TicketId),
			CreatedAt: helper.GetTimePointer(booking.CreatedAt.AsTime()),
			UpdatedAt: helper.GetTimePointer(booking.UpdatedAt.AsTime()),
		}
		bookingRes = append(bookingRes, bookingFullRes)
	}

	// graphql response
	resQ := &model.BookingResponse{
		Data:  bookingRes,
		Page:  helper.GetIntPointer(int(bookings.Page)),
		Limit: limit,
		Total: helper.GetIntPointer(int(bookings.Total)),
	}
	return resQ, nil
}

// Booking is the resolver for the Booking field.
func (r *mutationResolver) Booking(ctx context.Context) (*model.BookingOps, error) {
	return &model.BookingOps{}, nil
}

// Booking is the resolver for the Booking field.
func (r *queryResolver) Booking(ctx context.Context) (*model.BookingQuery, error) {
	return &model.BookingQuery{}, nil
}

// BookingOps returns graph.BookingOpsResolver implementation.
func (r *Resolver) BookingOps() graph.BookingOpsResolver { return &bookingOpsResolver{r} }

// BookingQuery returns graph.BookingQueryResolver implementation.
func (r *Resolver) BookingQuery() graph.BookingQueryResolver { return &bookingQueryResolver{r} }

type bookingOpsResolver struct{ *Resolver }
type bookingQueryResolver struct{ *Resolver }
