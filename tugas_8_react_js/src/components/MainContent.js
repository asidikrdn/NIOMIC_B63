import React, { useState } from "react";
import {
  Container,
  Carousel,
  Stack,
  ListGroup,
  Form,
  Button,
  Figure,
  Modal,
} from "react-bootstrap";

const MainContent = () => {
  const [show, setShow] = useState(false);

  const handleClose = () => setShow(false);
  const handleShow = (e) => {
    e.preventDefault();
    setShow(true);
  };

  return (
    <Container className="py-3 px-0" fluid>
      {/* Modals */}
      <Modal show={show} onHide={handleClose}>
        <Modal.Header closeButton>
          <Modal.Title>Login</Modal.Title>
        </Modal.Header>
        <Modal.Body>Anda Berhasil Login</Modal.Body>
        <Modal.Footer>
          <Button variant="secondary" onClick={handleClose}>
            Close
          </Button>
          <Button variant="primary" onClick={handleClose}>
            Save Username & Password
          </Button>
        </Modal.Footer>
      </Modal>
      {/* End of Modals */}

      {/* Carousel */}
      <Carousel className="mx-auto" style={{ width: "75%" }}>
        <Carousel.Item>
          <img
            className="d-block w-100"
            src="https://images.performgroup.com/di/library/GOAL/63/7f/sani-rizki-fauzi-indonesia-u-22_4yztc4nehkt61fa7omuyagf71.jpg?t=351764441&quality=60&w=1600"
            alt="First slide"
          />
          <Carousel.Caption>
            <h3
              className="fw-bold"
              style={{ textShadow: "0px 0px 10px black" }}
            >
              Sani Rizki Fauzi Tegaskan Mental Pemain Timnas Indonesia U-23
            </h3>
          </Carousel.Caption>
        </Carousel.Item>
        <Carousel.Item>
          <img
            className="d-block w-100"
            src="https://images.performgroup.com/di/library/GOAL/30/f0/neymar-psg-paris-saint-germain-2018-19_gmakk5891n9c19udcwofmrsgc.jpg?t=-1258742791&quality=60&w=1600"
            alt="First slide"
          />
          <Carousel.Caption>
            <h3
              className="fw-bold"
              style={{ textShadow: "0px 0px 10px black" }}
            >
              Neymar Bahas Kontrak Baru Di Paris Saint-Germaint
            </h3>
          </Carousel.Caption>
        </Carousel.Item>
        <Carousel.Item>
          <img
            className="d-block w-100"
            src="https://images.performgroup.com/di/library/GOAL/f3/33/england-celebrate-vs-montenegro_e97kv5anpqlq1gr9nc2qaz4ko.jpg?t=-1232521935&quality=60&w=1600"
            alt="First slide"
          />
          <Carousel.Caption>
            <h3
              className="fw-bold"
              style={{ textShadow: "0px 0px 10px black" }}
            >
              Inggris Mau Jadi Tim Terbaik Di Dunia
            </h3>
          </Carousel.Caption>
        </Carousel.Item>
      </Carousel>
      {/* End of Carousel */}

      <Stack
        direction="horizontal"
        gap={1}
        className="justify-content-between align-items-start mt-4"
      >
        {/* List Gorup */}
        <ListGroup style={{ width: "30%" }}>
          <ListGroup.Item active>Liga 1 Indonesia</ListGroup.Item>
          <ListGroup.Item>Liga Primer Inggris</ListGroup.Item>
          <ListGroup.Item>Divisi Primera</ListGroup.Item>
          <ListGroup.Item>Serie A</ListGroup.Item>
          <ListGroup.Item>League 1</ListGroup.Item>
          <ListGroup.Item>Bundles Liga</ListGroup.Item>
        </ListGroup>
        {/* End of List Gorup */}

        {/* Jumbotron */}
        {/* Karena jumbotron di Bootstrap 5 sudah tidak ada, sebagai gantinya saya buat manual ya kak */}
        <div
          className="px-4 py-5 bg-secondary bg-opacity-25 text-black rounded"
          style={{ width: "38%" }}
        >
          <h1 className="fw-semibold">DIVISI PRIMERA</h1>
          <p className="lead fs-6">
            Main Untuk Catalunya, Gerard Pique Minta Dihormati
          </p>
          <Button variant="primary">Read More</Button>
        </div>
        {/* End of Jumbotron */}

        <div style={{ width: "30%" }} className="pe-2">
          {/* Form */}

          <Figure className="text-center w-100">
            <Figure.Image
              width={75}
              height={180}
              alt="Silahkan Login"
              src="https://www.kindpng.com/picc/m/237-2370268_blue-register-icon-png-transparent-png.png"
            />
            <Figure.Caption className="text-black text-center fw-semibold fs-2">
              Silahkan Login
            </Figure.Caption>
          </Figure>

          <Form>
            <Form.Group className="mb-3" controlId="formBasicEmail">
              <Form.Control type="email" placeholder="Masukkan email" />
            </Form.Group>

            <Form.Group className="mb-3" controlId="formBasicPassword">
              <Form.Control type="password" placeholder="Password" />
            </Form.Group>
            <Form.Group className="mb-3" controlId="formBasicCheckbox">
              <Form.Check type="checkbox" label="Terms & Condition" />
            </Form.Group>
            <Button variant="primary" type="submit" onClick={handleShow}>
              Login
            </Button>
          </Form>
          {/* End of Form */}
        </div>
      </Stack>
    </Container>
  );
};

export default MainContent;
