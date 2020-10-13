# DB_RAYER_01
A Hetic student project.

## Usage

Required development environment:
- [Docker](https://www.docker.com)
- [Docker Compose](https://docs.docker.com/compose/install/)

Configure the development environment on your local machine:
```bash
$ git clone <this repo>
$ cd DB_RAYER_01
$ docker-compose up
```

You can now access the api: [http://localhost:1323/](http://localhost:1323/)

## Endpoints

```bash
## Get all customers and infos
GET(api/customer/id)

## Get all products of an order
GET(api/order/:id/product)

## Get all employees
GET(api/employee)

## Get all offices
GET(api/office)
```

## Documentation

The technical documentation of the project is available on GoDoc here : [Technical Documentation](https://godoc.org/github.com/HETIC-MT-P2021/DB_RAYER_P01/api)

## Author

<table>
  <tr>
    <td align="center">
      <a href="https://github.com/Akecel">
        <img src="https://github.com/Akecel.png" width="150px;"/><br>
        <b>Axel Rayer</b>
      </a>
    </td>
  </tr>
</table>


