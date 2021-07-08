package postgres

// Code generated by xo. DO NOT EDIT.

import (
	"context"
	"database/sql"
)

// Employee represents a row from 'public.employees'.
type Employee struct {
	EmployeeID      int            `json:"employee_id"`       // employee_id
	LastName        string         `json:"last_name"`         // last_name
	FirstName       string         `json:"first_name"`        // first_name
	Title           sql.NullString `json:"title"`             // title
	TitleOfCourtesy sql.NullString `json:"title_of_courtesy"` // title_of_courtesy
	BirthDate       sql.NullTime   `json:"birth_date"`        // birth_date
	HireDate        sql.NullTime   `json:"hire_date"`         // hire_date
	Address         sql.NullString `json:"address"`           // address
	City            sql.NullString `json:"city"`              // city
	Region          sql.NullString `json:"region"`            // region
	PostalCode      sql.NullString `json:"postal_code"`       // postal_code
	Country         sql.NullString `json:"country"`           // country
	HomePhone       sql.NullString `json:"home_phone"`        // home_phone
	Extension       sql.NullString `json:"extension"`         // extension
	Photo           []byte         `json:"photo"`             // photo
	Notes           sql.NullString `json:"notes"`             // notes
	ReportsTo       sql.NullInt64  `json:"reports_to"`        // reports_to
	PhotoPath       sql.NullString `json:"photo_path"`        // photo_path
	// xo fields
	_exists, _deleted bool
}

// Exists returns true when the Employee exists in the database.
func (e *Employee) Exists() bool {
	return e._exists
}

// Deleted returns true when the Employee has been marked for deletion from
// the database.
func (e *Employee) Deleted() bool {
	return e._deleted
}

// Insert inserts the Employee to the database.
func (e *Employee) Insert(ctx context.Context, db DB) error {
	switch {
	case e._exists: // already exists
		return logerror(&ErrInsertFailed{ErrAlreadyExists})
	case e._deleted: // deleted
		return logerror(&ErrInsertFailed{ErrMarkedForDeletion})
	}
	// insert (manual)
	const sqlstr = `INSERT INTO public.employees (` +
		`employee_id, last_name, first_name, title, title_of_courtesy, birth_date, hire_date, address, city, region, postal_code, country, home_phone, extension, photo, notes, reports_to, photo_path` +
		`) VALUES (` +
		`$1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15, $16, $17, $18` +
		`)`
	// run
	logf(sqlstr, e.EmployeeID, e.LastName, e.FirstName, e.Title, e.TitleOfCourtesy, e.BirthDate, e.HireDate, e.Address, e.City, e.Region, e.PostalCode, e.Country, e.HomePhone, e.Extension, e.Photo, e.Notes, e.ReportsTo, e.PhotoPath)
	if _, err := db.ExecContext(ctx, sqlstr, e.EmployeeID, e.LastName, e.FirstName, e.Title, e.TitleOfCourtesy, e.BirthDate, e.HireDate, e.Address, e.City, e.Region, e.PostalCode, e.Country, e.HomePhone, e.Extension, e.Photo, e.Notes, e.ReportsTo, e.PhotoPath); err != nil {
		return logerror(err)
	}
	// set exists
	e._exists = true
	return nil
}

// Update updates a Employee in the database.
func (e *Employee) Update(ctx context.Context, db DB) error {
	switch {
	case !e._exists: // doesn't exist
		return logerror(&ErrUpdateFailed{ErrDoesNotExist})
	case e._deleted: // deleted
		return logerror(&ErrUpdateFailed{ErrMarkedForDeletion})
	}
	// update with composite primary key
	const sqlstr = `UPDATE public.employees SET ` +
		`last_name = $1, first_name = $2, title = $3, title_of_courtesy = $4, birth_date = $5, hire_date = $6, address = $7, city = $8, region = $9, postal_code = $10, country = $11, home_phone = $12, extension = $13, photo = $14, notes = $15, reports_to = $16, photo_path = $17 ` +
		`WHERE employee_id = $18`
	// run
	logf(sqlstr, e.LastName, e.FirstName, e.Title, e.TitleOfCourtesy, e.BirthDate, e.HireDate, e.Address, e.City, e.Region, e.PostalCode, e.Country, e.HomePhone, e.Extension, e.Photo, e.Notes, e.ReportsTo, e.PhotoPath, e.EmployeeID)
	if _, err := db.ExecContext(ctx, sqlstr, e.LastName, e.FirstName, e.Title, e.TitleOfCourtesy, e.BirthDate, e.HireDate, e.Address, e.City, e.Region, e.PostalCode, e.Country, e.HomePhone, e.Extension, e.Photo, e.Notes, e.ReportsTo, e.PhotoPath, e.EmployeeID); err != nil {
		return logerror(err)
	}
	return nil
}

// Save saves the Employee to the database.
func (e *Employee) Save(ctx context.Context, db DB) error {
	if e.Exists() {
		return e.Update(ctx, db)
	}
	return e.Insert(ctx, db)
}

// Upsert performs an upsert for Employee.
func (e *Employee) Upsert(ctx context.Context, db DB) error {
	switch {
	case e._deleted: // deleted
		return logerror(&ErrUpsertFailed{ErrMarkedForDeletion})
	}
	// upsert
	const sqlstr = `INSERT INTO public.employees (` +
		`employee_id, last_name, first_name, title, title_of_courtesy, birth_date, hire_date, address, city, region, postal_code, country, home_phone, extension, photo, notes, reports_to, photo_path` +
		`) VALUES (` +
		`$1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15, $16, $17, $18` +
		`)` +
		` ON CONFLICT (employee_id) DO ` +
		`UPDATE SET ` +
		`last_name = EXCLUDED.last_name, first_name = EXCLUDED.first_name, title = EXCLUDED.title, title_of_courtesy = EXCLUDED.title_of_courtesy, birth_date = EXCLUDED.birth_date, hire_date = EXCLUDED.hire_date, address = EXCLUDED.address, city = EXCLUDED.city, region = EXCLUDED.region, postal_code = EXCLUDED.postal_code, country = EXCLUDED.country, home_phone = EXCLUDED.home_phone, extension = EXCLUDED.extension, photo = EXCLUDED.photo, notes = EXCLUDED.notes, reports_to = EXCLUDED.reports_to, photo_path = EXCLUDED.photo_path `
	// run
	logf(sqlstr, e.EmployeeID, e.LastName, e.FirstName, e.Title, e.TitleOfCourtesy, e.BirthDate, e.HireDate, e.Address, e.City, e.Region, e.PostalCode, e.Country, e.HomePhone, e.Extension, e.Photo, e.Notes, e.ReportsTo, e.PhotoPath)
	if _, err := db.ExecContext(ctx, sqlstr, e.EmployeeID, e.LastName, e.FirstName, e.Title, e.TitleOfCourtesy, e.BirthDate, e.HireDate, e.Address, e.City, e.Region, e.PostalCode, e.Country, e.HomePhone, e.Extension, e.Photo, e.Notes, e.ReportsTo, e.PhotoPath); err != nil {
		return err
	}
	// set exists
	e._exists = true
	return nil
}

// Delete deletes the Employee from the database.
func (e *Employee) Delete(ctx context.Context, db DB) error {
	switch {
	case !e._exists: // doesn't exist
		return nil
	case e._deleted: // deleted
		return nil
	}
	// delete with single primary key
	const sqlstr = `DELETE FROM public.employees ` +
		`WHERE employee_id = $1`
	// run
	logf(sqlstr, e.EmployeeID)
	if _, err := db.ExecContext(ctx, sqlstr, e.EmployeeID); err != nil {
		return logerror(err)
	}
	// set deleted
	e._deleted = true
	return nil
}

// EmployeeByEmployeeID retrieves a row from 'public.employees' as a Employee.
//
// Generated from index 'employees_pkey'.
func EmployeeByEmployeeID(ctx context.Context, db DB, employeeID int) (*Employee, error) {
	// query
	const sqlstr = `SELECT ` +
		`employee_id, last_name, first_name, title, title_of_courtesy, birth_date, hire_date, address, city, region, postal_code, country, home_phone, extension, photo, notes, reports_to, photo_path ` +
		`FROM public.employees ` +
		`WHERE employee_id = $1`
	// run
	logf(sqlstr, employeeID)
	e := Employee{
		_exists: true,
	}
	if err := db.QueryRowContext(ctx, sqlstr, employeeID).Scan(&e.EmployeeID, &e.LastName, &e.FirstName, &e.Title, &e.TitleOfCourtesy, &e.BirthDate, &e.HireDate, &e.Address, &e.City, &e.Region, &e.PostalCode, &e.Country, &e.HomePhone, &e.Extension, &e.Photo, &e.Notes, &e.ReportsTo, &e.PhotoPath); err != nil {
		return nil, logerror(err)
	}
	return &e, nil
}

// Employee returns the Employee associated with the Employee's (ReportsTo).
//
// Generated from foreign key 'employees_reports_to_fkey'.
func (e *Employee) Employee(ctx context.Context, db DB) (*Employee, error) {
	return EmployeeByEmployeeID(ctx, db, int(e.ReportsTo.Int64))
}
