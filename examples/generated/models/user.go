package models

import "github.com/go-swagger/go-swagger/strfmt"

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

/*User user

swagger:model User
*/
type User struct {

	/* Email email
	 */
	Email string `json:"email,omitempty"`

	/* FirstName first name
	 */
	FirstName string `json:"firstName,omitempty"`

	/* ID id
	 */
	ID int64 `json:"id,omitempty"`

	/* LastName last name
	 */
	LastName string `json:"lastName,omitempty"`

	/* Password password
	 */
	Password string `json:"password,omitempty"`

	/* Phone phone
	 */
	Phone string `json:"phone,omitempty"`

	/* User Status
	 */
	UserStatus int32 `json:"userStatus,omitempty"`

	/* Username username
	 */
	Username string `json:"username,omitempty"`
}

// Validate validates this user
func (m *User) Validate(formats strfmt.Registry) error {
	return nil
}
