package sqlite3

// Code generated by xo. DO NOT EDIT.

import (
	"context"
)

// Territory represents a row from 'territories'.
type Territory struct {
	TerritoryID          string `json:"territory_id"`          // territory_id
	TerritoryDescription string `json:"territory_description"` // territory_description
	RegionID             int    `json:"region_id"`             // region_id
	// xo fields
	_exists, _deleted bool
}

// Exists returns true when the Territory exists in the database.
func (t *Territory) Exists() bool {
	return t._exists
}

// Deleted returns true when the Territory has been marked for deletion from
// the database.
func (t *Territory) Deleted() bool {
	return t._deleted
}

// Insert inserts the Territory to the database.
func (t *Territory) Insert(ctx context.Context, db DB) error {
	switch {
	case t._exists: // already exists
		return logerror(&ErrInsertFailed{ErrAlreadyExists})
	case t._deleted: // deleted
		return logerror(&ErrInsertFailed{ErrMarkedForDeletion})
	}
	// insert (manual)
	const sqlstr = `INSERT INTO territories (` +
		`territory_id, territory_description, region_id` +
		`) VALUES (` +
		`$1, $2, $3` +
		`)`
	// run
	logf(sqlstr, t.TerritoryID, t.TerritoryDescription, t.RegionID)
	if _, err := db.ExecContext(ctx, sqlstr, t.TerritoryID, t.TerritoryDescription, t.RegionID); err != nil {
		return logerror(err)
	}
	// set exists
	t._exists = true
	return nil
}

// Update updates a Territory in the database.
func (t *Territory) Update(ctx context.Context, db DB) error {
	switch {
	case !t._exists: // doesn't exist
		return logerror(&ErrUpdateFailed{ErrDoesNotExist})
	case t._deleted: // deleted
		return logerror(&ErrUpdateFailed{ErrMarkedForDeletion})
	}
	// update with primary key
	const sqlstr = `UPDATE territories SET ` +
		`territory_description = $1, region_id = $2 ` +
		`WHERE territory_id = $3`
	// run
	logf(sqlstr, t.TerritoryDescription, t.RegionID, t.TerritoryID)
	if _, err := db.ExecContext(ctx, sqlstr, t.TerritoryDescription, t.RegionID, t.TerritoryID); err != nil {
		return logerror(err)
	}
	return nil
}

// Save saves the Territory to the database.
func (t *Territory) Save(ctx context.Context, db DB) error {
	if t.Exists() {
		return t.Update(ctx, db)
	}
	return t.Insert(ctx, db)
}

// Upsert performs an upsert for Territory.
func (t *Territory) Upsert(ctx context.Context, db DB) error {
	switch {
	case t._deleted: // deleted
		return logerror(&ErrUpsertFailed{ErrMarkedForDeletion})
	}
	// upsert
	const sqlstr = `INSERT INTO territories (` +
		`territory_id, territory_description, region_id` +
		`) VALUES (` +
		`$1, $2, $3` +
		`)` +
		` ON CONFLICT (territory_id) DO ` +
		`UPDATE SET ` +
		`territory_description = EXCLUDED.territory_description, region_id = EXCLUDED.region_id `
	// run
	logf(sqlstr, t.TerritoryID, t.TerritoryDescription, t.RegionID)
	if _, err := db.ExecContext(ctx, sqlstr, t.TerritoryID, t.TerritoryDescription, t.RegionID); err != nil {
		return err
	}
	// set exists
	t._exists = true
	return nil
}

// Delete deletes the Territory from the database.
func (t *Territory) Delete(ctx context.Context, db DB) error {
	switch {
	case !t._exists: // doesn't exist
		return nil
	case t._deleted: // deleted
		return nil
	}
	// delete with single primary key
	const sqlstr = `DELETE FROM territories ` +
		`WHERE territory_id = $1`
	// run
	logf(sqlstr, t.TerritoryID)
	if _, err := db.ExecContext(ctx, sqlstr, t.TerritoryID); err != nil {
		return logerror(err)
	}
	// set deleted
	t._deleted = true
	return nil
}

// TerritoryByTerritoryID retrieves a row from 'territories' as a Territory.
//
// Generated from index 'sqlite_autoindex_territories_1'.
func TerritoryByTerritoryID(ctx context.Context, db DB, territoryID string) (*Territory, error) {
	// query
	const sqlstr = `SELECT ` +
		`territory_id, territory_description, region_id ` +
		`FROM territories ` +
		`WHERE territory_id = $1`
	// run
	logf(sqlstr, territoryID)
	t := Territory{
		_exists: true,
	}
	if err := db.QueryRowContext(ctx, sqlstr, territoryID).Scan(&t.TerritoryID, &t.TerritoryDescription, &t.RegionID); err != nil {
		return nil, logerror(err)
	}
	return &t, nil
}

// Region returns the Region associated with the Territory's (RegionID).
//
// Generated from foreign key 'territories_region_id_fkey'.
func (t *Territory) Region(ctx context.Context, db DB) (*Region, error) {
	return RegionByRegionID(ctx, db, t.RegionID)
}
