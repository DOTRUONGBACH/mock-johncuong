// Code generated by ent, DO NOT EDIT.

package customer

import (
	"mock-project/grpc/customer-grpc/ent/predicate"
	"time"

	"entgo.io/ent/dialect/sql"
)

// ID filters vertices based on their ID field.
func ID(id int64) predicate.Customer {
	return predicate.Customer(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int64) predicate.Customer {
	return predicate.Customer(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int64) predicate.Customer {
	return predicate.Customer(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int64) predicate.Customer {
	return predicate.Customer(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int64) predicate.Customer {
	return predicate.Customer(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int64) predicate.Customer {
	return predicate.Customer(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int64) predicate.Customer {
	return predicate.Customer(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int64) predicate.Customer {
	return predicate.Customer(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int64) predicate.Customer {
	return predicate.Customer(sql.FieldLTE(FieldID, id))
}

// Name applies equality check predicate on the "name" field. It's identical to NameEQ.
func Name(v string) predicate.Customer {
	return predicate.Customer(sql.FieldEQ(FieldName, v))
}

// Email applies equality check predicate on the "email" field. It's identical to EmailEQ.
func Email(v string) predicate.Customer {
	return predicate.Customer(sql.FieldEQ(FieldEmail, v))
}

// Address applies equality check predicate on the "address" field. It's identical to AddressEQ.
func Address(v string) predicate.Customer {
	return predicate.Customer(sql.FieldEQ(FieldAddress, v))
}

// PhoneNumber applies equality check predicate on the "phone_number" field. It's identical to PhoneNumberEQ.
func PhoneNumber(v string) predicate.Customer {
	return predicate.Customer(sql.FieldEQ(FieldPhoneNumber, v))
}

// IdentifyNumber applies equality check predicate on the "identify_number" field. It's identical to IdentifyNumberEQ.
func IdentifyNumber(v string) predicate.Customer {
	return predicate.Customer(sql.FieldEQ(FieldIdentifyNumber, v))
}

// DateOfBirth applies equality check predicate on the "date_of_birth" field. It's identical to DateOfBirthEQ.
func DateOfBirth(v time.Time) predicate.Customer {
	return predicate.Customer(sql.FieldEQ(FieldDateOfBirth, v))
}

// MemberCode applies equality check predicate on the "member_code" field. It's identical to MemberCodeEQ.
func MemberCode(v string) predicate.Customer {
	return predicate.Customer(sql.FieldEQ(FieldMemberCode, v))
}

// CreatedAt applies equality check predicate on the "created_at" field. It's identical to CreatedAtEQ.
func CreatedAt(v time.Time) predicate.Customer {
	return predicate.Customer(sql.FieldEQ(FieldCreatedAt, v))
}

// UpdatedAt applies equality check predicate on the "updated_at" field. It's identical to UpdatedAtEQ.
func UpdatedAt(v time.Time) predicate.Customer {
	return predicate.Customer(sql.FieldEQ(FieldUpdatedAt, v))
}

// NameEQ applies the EQ predicate on the "name" field.
func NameEQ(v string) predicate.Customer {
	return predicate.Customer(sql.FieldEQ(FieldName, v))
}

// NameNEQ applies the NEQ predicate on the "name" field.
func NameNEQ(v string) predicate.Customer {
	return predicate.Customer(sql.FieldNEQ(FieldName, v))
}

// NameIn applies the In predicate on the "name" field.
func NameIn(vs ...string) predicate.Customer {
	return predicate.Customer(sql.FieldIn(FieldName, vs...))
}

// NameNotIn applies the NotIn predicate on the "name" field.
func NameNotIn(vs ...string) predicate.Customer {
	return predicate.Customer(sql.FieldNotIn(FieldName, vs...))
}

// NameGT applies the GT predicate on the "name" field.
func NameGT(v string) predicate.Customer {
	return predicate.Customer(sql.FieldGT(FieldName, v))
}

// NameGTE applies the GTE predicate on the "name" field.
func NameGTE(v string) predicate.Customer {
	return predicate.Customer(sql.FieldGTE(FieldName, v))
}

// NameLT applies the LT predicate on the "name" field.
func NameLT(v string) predicate.Customer {
	return predicate.Customer(sql.FieldLT(FieldName, v))
}

// NameLTE applies the LTE predicate on the "name" field.
func NameLTE(v string) predicate.Customer {
	return predicate.Customer(sql.FieldLTE(FieldName, v))
}

// NameContains applies the Contains predicate on the "name" field.
func NameContains(v string) predicate.Customer {
	return predicate.Customer(sql.FieldContains(FieldName, v))
}

// NameHasPrefix applies the HasPrefix predicate on the "name" field.
func NameHasPrefix(v string) predicate.Customer {
	return predicate.Customer(sql.FieldHasPrefix(FieldName, v))
}

// NameHasSuffix applies the HasSuffix predicate on the "name" field.
func NameHasSuffix(v string) predicate.Customer {
	return predicate.Customer(sql.FieldHasSuffix(FieldName, v))
}

// NameEqualFold applies the EqualFold predicate on the "name" field.
func NameEqualFold(v string) predicate.Customer {
	return predicate.Customer(sql.FieldEqualFold(FieldName, v))
}

// NameContainsFold applies the ContainsFold predicate on the "name" field.
func NameContainsFold(v string) predicate.Customer {
	return predicate.Customer(sql.FieldContainsFold(FieldName, v))
}

// EmailEQ applies the EQ predicate on the "email" field.
func EmailEQ(v string) predicate.Customer {
	return predicate.Customer(sql.FieldEQ(FieldEmail, v))
}

// EmailNEQ applies the NEQ predicate on the "email" field.
func EmailNEQ(v string) predicate.Customer {
	return predicate.Customer(sql.FieldNEQ(FieldEmail, v))
}

// EmailIn applies the In predicate on the "email" field.
func EmailIn(vs ...string) predicate.Customer {
	return predicate.Customer(sql.FieldIn(FieldEmail, vs...))
}

// EmailNotIn applies the NotIn predicate on the "email" field.
func EmailNotIn(vs ...string) predicate.Customer {
	return predicate.Customer(sql.FieldNotIn(FieldEmail, vs...))
}

// EmailGT applies the GT predicate on the "email" field.
func EmailGT(v string) predicate.Customer {
	return predicate.Customer(sql.FieldGT(FieldEmail, v))
}

// EmailGTE applies the GTE predicate on the "email" field.
func EmailGTE(v string) predicate.Customer {
	return predicate.Customer(sql.FieldGTE(FieldEmail, v))
}

// EmailLT applies the LT predicate on the "email" field.
func EmailLT(v string) predicate.Customer {
	return predicate.Customer(sql.FieldLT(FieldEmail, v))
}

// EmailLTE applies the LTE predicate on the "email" field.
func EmailLTE(v string) predicate.Customer {
	return predicate.Customer(sql.FieldLTE(FieldEmail, v))
}

// EmailContains applies the Contains predicate on the "email" field.
func EmailContains(v string) predicate.Customer {
	return predicate.Customer(sql.FieldContains(FieldEmail, v))
}

// EmailHasPrefix applies the HasPrefix predicate on the "email" field.
func EmailHasPrefix(v string) predicate.Customer {
	return predicate.Customer(sql.FieldHasPrefix(FieldEmail, v))
}

// EmailHasSuffix applies the HasSuffix predicate on the "email" field.
func EmailHasSuffix(v string) predicate.Customer {
	return predicate.Customer(sql.FieldHasSuffix(FieldEmail, v))
}

// EmailEqualFold applies the EqualFold predicate on the "email" field.
func EmailEqualFold(v string) predicate.Customer {
	return predicate.Customer(sql.FieldEqualFold(FieldEmail, v))
}

// EmailContainsFold applies the ContainsFold predicate on the "email" field.
func EmailContainsFold(v string) predicate.Customer {
	return predicate.Customer(sql.FieldContainsFold(FieldEmail, v))
}

// AddressEQ applies the EQ predicate on the "address" field.
func AddressEQ(v string) predicate.Customer {
	return predicate.Customer(sql.FieldEQ(FieldAddress, v))
}

// AddressNEQ applies the NEQ predicate on the "address" field.
func AddressNEQ(v string) predicate.Customer {
	return predicate.Customer(sql.FieldNEQ(FieldAddress, v))
}

// AddressIn applies the In predicate on the "address" field.
func AddressIn(vs ...string) predicate.Customer {
	return predicate.Customer(sql.FieldIn(FieldAddress, vs...))
}

// AddressNotIn applies the NotIn predicate on the "address" field.
func AddressNotIn(vs ...string) predicate.Customer {
	return predicate.Customer(sql.FieldNotIn(FieldAddress, vs...))
}

// AddressGT applies the GT predicate on the "address" field.
func AddressGT(v string) predicate.Customer {
	return predicate.Customer(sql.FieldGT(FieldAddress, v))
}

// AddressGTE applies the GTE predicate on the "address" field.
func AddressGTE(v string) predicate.Customer {
	return predicate.Customer(sql.FieldGTE(FieldAddress, v))
}

// AddressLT applies the LT predicate on the "address" field.
func AddressLT(v string) predicate.Customer {
	return predicate.Customer(sql.FieldLT(FieldAddress, v))
}

// AddressLTE applies the LTE predicate on the "address" field.
func AddressLTE(v string) predicate.Customer {
	return predicate.Customer(sql.FieldLTE(FieldAddress, v))
}

// AddressContains applies the Contains predicate on the "address" field.
func AddressContains(v string) predicate.Customer {
	return predicate.Customer(sql.FieldContains(FieldAddress, v))
}

// AddressHasPrefix applies the HasPrefix predicate on the "address" field.
func AddressHasPrefix(v string) predicate.Customer {
	return predicate.Customer(sql.FieldHasPrefix(FieldAddress, v))
}

// AddressHasSuffix applies the HasSuffix predicate on the "address" field.
func AddressHasSuffix(v string) predicate.Customer {
	return predicate.Customer(sql.FieldHasSuffix(FieldAddress, v))
}

// AddressEqualFold applies the EqualFold predicate on the "address" field.
func AddressEqualFold(v string) predicate.Customer {
	return predicate.Customer(sql.FieldEqualFold(FieldAddress, v))
}

// AddressContainsFold applies the ContainsFold predicate on the "address" field.
func AddressContainsFold(v string) predicate.Customer {
	return predicate.Customer(sql.FieldContainsFold(FieldAddress, v))
}

// PhoneNumberEQ applies the EQ predicate on the "phone_number" field.
func PhoneNumberEQ(v string) predicate.Customer {
	return predicate.Customer(sql.FieldEQ(FieldPhoneNumber, v))
}

// PhoneNumberNEQ applies the NEQ predicate on the "phone_number" field.
func PhoneNumberNEQ(v string) predicate.Customer {
	return predicate.Customer(sql.FieldNEQ(FieldPhoneNumber, v))
}

// PhoneNumberIn applies the In predicate on the "phone_number" field.
func PhoneNumberIn(vs ...string) predicate.Customer {
	return predicate.Customer(sql.FieldIn(FieldPhoneNumber, vs...))
}

// PhoneNumberNotIn applies the NotIn predicate on the "phone_number" field.
func PhoneNumberNotIn(vs ...string) predicate.Customer {
	return predicate.Customer(sql.FieldNotIn(FieldPhoneNumber, vs...))
}

// PhoneNumberGT applies the GT predicate on the "phone_number" field.
func PhoneNumberGT(v string) predicate.Customer {
	return predicate.Customer(sql.FieldGT(FieldPhoneNumber, v))
}

// PhoneNumberGTE applies the GTE predicate on the "phone_number" field.
func PhoneNumberGTE(v string) predicate.Customer {
	return predicate.Customer(sql.FieldGTE(FieldPhoneNumber, v))
}

// PhoneNumberLT applies the LT predicate on the "phone_number" field.
func PhoneNumberLT(v string) predicate.Customer {
	return predicate.Customer(sql.FieldLT(FieldPhoneNumber, v))
}

// PhoneNumberLTE applies the LTE predicate on the "phone_number" field.
func PhoneNumberLTE(v string) predicate.Customer {
	return predicate.Customer(sql.FieldLTE(FieldPhoneNumber, v))
}

// PhoneNumberContains applies the Contains predicate on the "phone_number" field.
func PhoneNumberContains(v string) predicate.Customer {
	return predicate.Customer(sql.FieldContains(FieldPhoneNumber, v))
}

// PhoneNumberHasPrefix applies the HasPrefix predicate on the "phone_number" field.
func PhoneNumberHasPrefix(v string) predicate.Customer {
	return predicate.Customer(sql.FieldHasPrefix(FieldPhoneNumber, v))
}

// PhoneNumberHasSuffix applies the HasSuffix predicate on the "phone_number" field.
func PhoneNumberHasSuffix(v string) predicate.Customer {
	return predicate.Customer(sql.FieldHasSuffix(FieldPhoneNumber, v))
}

// PhoneNumberEqualFold applies the EqualFold predicate on the "phone_number" field.
func PhoneNumberEqualFold(v string) predicate.Customer {
	return predicate.Customer(sql.FieldEqualFold(FieldPhoneNumber, v))
}

// PhoneNumberContainsFold applies the ContainsFold predicate on the "phone_number" field.
func PhoneNumberContainsFold(v string) predicate.Customer {
	return predicate.Customer(sql.FieldContainsFold(FieldPhoneNumber, v))
}

// IdentifyNumberEQ applies the EQ predicate on the "identify_number" field.
func IdentifyNumberEQ(v string) predicate.Customer {
	return predicate.Customer(sql.FieldEQ(FieldIdentifyNumber, v))
}

// IdentifyNumberNEQ applies the NEQ predicate on the "identify_number" field.
func IdentifyNumberNEQ(v string) predicate.Customer {
	return predicate.Customer(sql.FieldNEQ(FieldIdentifyNumber, v))
}

// IdentifyNumberIn applies the In predicate on the "identify_number" field.
func IdentifyNumberIn(vs ...string) predicate.Customer {
	return predicate.Customer(sql.FieldIn(FieldIdentifyNumber, vs...))
}

// IdentifyNumberNotIn applies the NotIn predicate on the "identify_number" field.
func IdentifyNumberNotIn(vs ...string) predicate.Customer {
	return predicate.Customer(sql.FieldNotIn(FieldIdentifyNumber, vs...))
}

// IdentifyNumberGT applies the GT predicate on the "identify_number" field.
func IdentifyNumberGT(v string) predicate.Customer {
	return predicate.Customer(sql.FieldGT(FieldIdentifyNumber, v))
}

// IdentifyNumberGTE applies the GTE predicate on the "identify_number" field.
func IdentifyNumberGTE(v string) predicate.Customer {
	return predicate.Customer(sql.FieldGTE(FieldIdentifyNumber, v))
}

// IdentifyNumberLT applies the LT predicate on the "identify_number" field.
func IdentifyNumberLT(v string) predicate.Customer {
	return predicate.Customer(sql.FieldLT(FieldIdentifyNumber, v))
}

// IdentifyNumberLTE applies the LTE predicate on the "identify_number" field.
func IdentifyNumberLTE(v string) predicate.Customer {
	return predicate.Customer(sql.FieldLTE(FieldIdentifyNumber, v))
}

// IdentifyNumberContains applies the Contains predicate on the "identify_number" field.
func IdentifyNumberContains(v string) predicate.Customer {
	return predicate.Customer(sql.FieldContains(FieldIdentifyNumber, v))
}

// IdentifyNumberHasPrefix applies the HasPrefix predicate on the "identify_number" field.
func IdentifyNumberHasPrefix(v string) predicate.Customer {
	return predicate.Customer(sql.FieldHasPrefix(FieldIdentifyNumber, v))
}

// IdentifyNumberHasSuffix applies the HasSuffix predicate on the "identify_number" field.
func IdentifyNumberHasSuffix(v string) predicate.Customer {
	return predicate.Customer(sql.FieldHasSuffix(FieldIdentifyNumber, v))
}

// IdentifyNumberEqualFold applies the EqualFold predicate on the "identify_number" field.
func IdentifyNumberEqualFold(v string) predicate.Customer {
	return predicate.Customer(sql.FieldEqualFold(FieldIdentifyNumber, v))
}

// IdentifyNumberContainsFold applies the ContainsFold predicate on the "identify_number" field.
func IdentifyNumberContainsFold(v string) predicate.Customer {
	return predicate.Customer(sql.FieldContainsFold(FieldIdentifyNumber, v))
}

// DateOfBirthEQ applies the EQ predicate on the "date_of_birth" field.
func DateOfBirthEQ(v time.Time) predicate.Customer {
	return predicate.Customer(sql.FieldEQ(FieldDateOfBirth, v))
}

// DateOfBirthNEQ applies the NEQ predicate on the "date_of_birth" field.
func DateOfBirthNEQ(v time.Time) predicate.Customer {
	return predicate.Customer(sql.FieldNEQ(FieldDateOfBirth, v))
}

// DateOfBirthIn applies the In predicate on the "date_of_birth" field.
func DateOfBirthIn(vs ...time.Time) predicate.Customer {
	return predicate.Customer(sql.FieldIn(FieldDateOfBirth, vs...))
}

// DateOfBirthNotIn applies the NotIn predicate on the "date_of_birth" field.
func DateOfBirthNotIn(vs ...time.Time) predicate.Customer {
	return predicate.Customer(sql.FieldNotIn(FieldDateOfBirth, vs...))
}

// DateOfBirthGT applies the GT predicate on the "date_of_birth" field.
func DateOfBirthGT(v time.Time) predicate.Customer {
	return predicate.Customer(sql.FieldGT(FieldDateOfBirth, v))
}

// DateOfBirthGTE applies the GTE predicate on the "date_of_birth" field.
func DateOfBirthGTE(v time.Time) predicate.Customer {
	return predicate.Customer(sql.FieldGTE(FieldDateOfBirth, v))
}

// DateOfBirthLT applies the LT predicate on the "date_of_birth" field.
func DateOfBirthLT(v time.Time) predicate.Customer {
	return predicate.Customer(sql.FieldLT(FieldDateOfBirth, v))
}

// DateOfBirthLTE applies the LTE predicate on the "date_of_birth" field.
func DateOfBirthLTE(v time.Time) predicate.Customer {
	return predicate.Customer(sql.FieldLTE(FieldDateOfBirth, v))
}

// MemberCodeEQ applies the EQ predicate on the "member_code" field.
func MemberCodeEQ(v string) predicate.Customer {
	return predicate.Customer(sql.FieldEQ(FieldMemberCode, v))
}

// MemberCodeNEQ applies the NEQ predicate on the "member_code" field.
func MemberCodeNEQ(v string) predicate.Customer {
	return predicate.Customer(sql.FieldNEQ(FieldMemberCode, v))
}

// MemberCodeIn applies the In predicate on the "member_code" field.
func MemberCodeIn(vs ...string) predicate.Customer {
	return predicate.Customer(sql.FieldIn(FieldMemberCode, vs...))
}

// MemberCodeNotIn applies the NotIn predicate on the "member_code" field.
func MemberCodeNotIn(vs ...string) predicate.Customer {
	return predicate.Customer(sql.FieldNotIn(FieldMemberCode, vs...))
}

// MemberCodeGT applies the GT predicate on the "member_code" field.
func MemberCodeGT(v string) predicate.Customer {
	return predicate.Customer(sql.FieldGT(FieldMemberCode, v))
}

// MemberCodeGTE applies the GTE predicate on the "member_code" field.
func MemberCodeGTE(v string) predicate.Customer {
	return predicate.Customer(sql.FieldGTE(FieldMemberCode, v))
}

// MemberCodeLT applies the LT predicate on the "member_code" field.
func MemberCodeLT(v string) predicate.Customer {
	return predicate.Customer(sql.FieldLT(FieldMemberCode, v))
}

// MemberCodeLTE applies the LTE predicate on the "member_code" field.
func MemberCodeLTE(v string) predicate.Customer {
	return predicate.Customer(sql.FieldLTE(FieldMemberCode, v))
}

// MemberCodeContains applies the Contains predicate on the "member_code" field.
func MemberCodeContains(v string) predicate.Customer {
	return predicate.Customer(sql.FieldContains(FieldMemberCode, v))
}

// MemberCodeHasPrefix applies the HasPrefix predicate on the "member_code" field.
func MemberCodeHasPrefix(v string) predicate.Customer {
	return predicate.Customer(sql.FieldHasPrefix(FieldMemberCode, v))
}

// MemberCodeHasSuffix applies the HasSuffix predicate on the "member_code" field.
func MemberCodeHasSuffix(v string) predicate.Customer {
	return predicate.Customer(sql.FieldHasSuffix(FieldMemberCode, v))
}

// MemberCodeIsNil applies the IsNil predicate on the "member_code" field.
func MemberCodeIsNil() predicate.Customer {
	return predicate.Customer(sql.FieldIsNull(FieldMemberCode))
}

// MemberCodeNotNil applies the NotNil predicate on the "member_code" field.
func MemberCodeNotNil() predicate.Customer {
	return predicate.Customer(sql.FieldNotNull(FieldMemberCode))
}

// MemberCodeEqualFold applies the EqualFold predicate on the "member_code" field.
func MemberCodeEqualFold(v string) predicate.Customer {
	return predicate.Customer(sql.FieldEqualFold(FieldMemberCode, v))
}

// MemberCodeContainsFold applies the ContainsFold predicate on the "member_code" field.
func MemberCodeContainsFold(v string) predicate.Customer {
	return predicate.Customer(sql.FieldContainsFold(FieldMemberCode, v))
}

// CreatedAtEQ applies the EQ predicate on the "created_at" field.
func CreatedAtEQ(v time.Time) predicate.Customer {
	return predicate.Customer(sql.FieldEQ(FieldCreatedAt, v))
}

// CreatedAtNEQ applies the NEQ predicate on the "created_at" field.
func CreatedAtNEQ(v time.Time) predicate.Customer {
	return predicate.Customer(sql.FieldNEQ(FieldCreatedAt, v))
}

// CreatedAtIn applies the In predicate on the "created_at" field.
func CreatedAtIn(vs ...time.Time) predicate.Customer {
	return predicate.Customer(sql.FieldIn(FieldCreatedAt, vs...))
}

// CreatedAtNotIn applies the NotIn predicate on the "created_at" field.
func CreatedAtNotIn(vs ...time.Time) predicate.Customer {
	return predicate.Customer(sql.FieldNotIn(FieldCreatedAt, vs...))
}

// CreatedAtGT applies the GT predicate on the "created_at" field.
func CreatedAtGT(v time.Time) predicate.Customer {
	return predicate.Customer(sql.FieldGT(FieldCreatedAt, v))
}

// CreatedAtGTE applies the GTE predicate on the "created_at" field.
func CreatedAtGTE(v time.Time) predicate.Customer {
	return predicate.Customer(sql.FieldGTE(FieldCreatedAt, v))
}

// CreatedAtLT applies the LT predicate on the "created_at" field.
func CreatedAtLT(v time.Time) predicate.Customer {
	return predicate.Customer(sql.FieldLT(FieldCreatedAt, v))
}

// CreatedAtLTE applies the LTE predicate on the "created_at" field.
func CreatedAtLTE(v time.Time) predicate.Customer {
	return predicate.Customer(sql.FieldLTE(FieldCreatedAt, v))
}

// UpdatedAtEQ applies the EQ predicate on the "updated_at" field.
func UpdatedAtEQ(v time.Time) predicate.Customer {
	return predicate.Customer(sql.FieldEQ(FieldUpdatedAt, v))
}

// UpdatedAtNEQ applies the NEQ predicate on the "updated_at" field.
func UpdatedAtNEQ(v time.Time) predicate.Customer {
	return predicate.Customer(sql.FieldNEQ(FieldUpdatedAt, v))
}

// UpdatedAtIn applies the In predicate on the "updated_at" field.
func UpdatedAtIn(vs ...time.Time) predicate.Customer {
	return predicate.Customer(sql.FieldIn(FieldUpdatedAt, vs...))
}

// UpdatedAtNotIn applies the NotIn predicate on the "updated_at" field.
func UpdatedAtNotIn(vs ...time.Time) predicate.Customer {
	return predicate.Customer(sql.FieldNotIn(FieldUpdatedAt, vs...))
}

// UpdatedAtGT applies the GT predicate on the "updated_at" field.
func UpdatedAtGT(v time.Time) predicate.Customer {
	return predicate.Customer(sql.FieldGT(FieldUpdatedAt, v))
}

// UpdatedAtGTE applies the GTE predicate on the "updated_at" field.
func UpdatedAtGTE(v time.Time) predicate.Customer {
	return predicate.Customer(sql.FieldGTE(FieldUpdatedAt, v))
}

// UpdatedAtLT applies the LT predicate on the "updated_at" field.
func UpdatedAtLT(v time.Time) predicate.Customer {
	return predicate.Customer(sql.FieldLT(FieldUpdatedAt, v))
}

// UpdatedAtLTE applies the LTE predicate on the "updated_at" field.
func UpdatedAtLTE(v time.Time) predicate.Customer {
	return predicate.Customer(sql.FieldLTE(FieldUpdatedAt, v))
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.Customer) predicate.Customer {
	return predicate.Customer(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.Customer) predicate.Customer {
	return predicate.Customer(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for i, p := range predicates {
			if i > 0 {
				s1.Or()
			}
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Not applies the not operator on the given predicate.
func Not(p predicate.Customer) predicate.Customer {
	return predicate.Customer(func(s *sql.Selector) {
		p(s.Not())
	})
}
