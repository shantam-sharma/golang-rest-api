if files are inside the same user folder they have same package | package users |

handler.go
Meaning:

read request
validate method
decode JSON
call service
send HTTP response

The handler should care about:

HTTP protocol
request/response lifecycle

NOT business rules.



Suppose JSON is a parcel.

Deserialization means:

open parcel
extract data
put data into Go struct


Response serialization
json.NewEncoder(w).Encode(user)

This is the REVERSE process.

Now Go converts struct → JSON.
