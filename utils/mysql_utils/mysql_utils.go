package mysql_utils

import (
	"strings"

	"github.com/gasparyanarman/bookstore_users-api/utils/errors"
	"github.com/go-sql-driver/mysql"
)

func ParseError(err error) *errors.RestErr {
	sqlErr, ok := err.(*mysql.MySQLError)

	if !ok {
		if strings.Contains(err.Error(), "no rows in result set") {
			return errors.NewNotFoundError("No record matching given id")
		}

		return errors.NewInternalServerError("Error parsing database response")
	}

	switch sqlErr.Number {
	case 1062:
		return errors.NewBadRequestError("Invalid Data")
	}

	return errors.NewInternalServerError("Error processing request")
}
