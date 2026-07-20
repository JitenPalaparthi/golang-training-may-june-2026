package models

import "errors"

type Product struct {
	// gorm.Model
	CommonModel
	Name        string  `json:"name" gorm:"size:200;not null"`
	Description string  `json:"description" gorm:"not null;default:some product"`
	Price       float32 `json:"price" gorm:"not null"`
	Stock       int     `json:"stock"`
	Status      string  `json:"status" gorm:"default:active"`
}

type CommonModel struct {
	ID uint `gorm:"primarykey"`
	// We explicitly cast the epoch to a bigint so Postgres drops the decimal point
	LastModified int64 `json:"last_modified" gorm:"column:last_modified;type:bigint;not null;default:extract(epoch from now())::bigint"`
}

func (p *Product) Validate() error {

	if p.Name == "" {
		return errors.New("invalid product name")
	}

	if p.Description == "" {
		return errors.New("invalid product description ")
	}

	if p.Price <= 0 {
		return errors.New("invalid product price.Price must be greater than 0")
	}

	return nil

}

/***
http 1.1

http2 protocol

mulplexing

Sending a request to a server using restful service

--> Client >>>> ree --> Server (TCP/IP connect  ^ http connection)
Client <<< rest <-- Server

CURD Operations
---------------
Bidirections Communication(WebSocket)
--------------------------
Server Side Event
-----------------
Server Pull/ Server Long pull

http2
------

A single tcp connection can handle multiple streams a.k.a

muliple operations can simultaniously/ concurrently happen on a single tcp/ip connection

stream(s)

Tcp >>>>>> Connection Server
| - - - - - - - - - - - -  | Unary
| - - - - - - - - - - - -  | Duplex
| - - - - - - - - - - - -  | Client Steam
| - - - - - - - - - - - -  | Server Stream
| - - - - - - - - - - - -  | Any number of streams
TCP <<<<<<<< Connection Close

Frame(s)

Http data transport format is generally test format (http1.1)
Content-Type --> application/json application/xml etc

http2 --> data transport in binary format

request
{
    "name":"Macbook Pro",
    "description": "Apple Macbook Pro M1 chip set",
    "price": 100123.12,
    "stock":100
}

response:

request
{	"id":"fab112321-02123123-abcdef123"
    "name":"Macbook Pro",
    "description": "Apple Macbook Pro M1 chip set",
    "price": 100123.12,
    "stock":100
	"status":"active"
}

This sending keys to the server and receiving data of keys from the server is already known why to transmit the keys/
Use some kind of compression which would serialize and deserialize at the end point - >> not possible using http1.1

This can be solved using a technique called hpack compression in http2

json
xml

In order to do encoading/decoading or serialization/deserialization http2 uses a format is called
protocol buffers

rpc

gRPC
protocol Buffer

The data that i want to send and receive is product and other forms
RPC I want to execure are Create,Get

- step -1 --> create a proto file
-

***/
