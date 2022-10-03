import {
  Container,
  Navbar,
  Nav,
  NavDropdown,
  Form,
  Button,
} from "react-bootstrap";

const NavigationBar = () => {
  return (
    <>
      <Navbar bg="dark" variant="dark" expand="lg">
        <Container fluid className="mx-0">
          <Navbar.Brand href="#home">Akses Sport</Navbar.Brand>
          <Navbar.Toggle aria-controls="basic-navbar-nav" />
          <Navbar.Collapse id="basic-navbar-nav">
            <Nav className="me-auto">
              <Nav.Link href="#home">Home</Nav.Link>
              <Nav.Link href="#berita">Berita</Nav.Link>
              <Nav.Link href="#live-scores">Live Scores</Nav.Link>
              <NavDropdown title="Piala & Liga" id="basic-nav-dropdown">
                <NavDropdown.Item href="#liga1">Liga 1</NavDropdown.Item>
                <NavDropdown.Item href="#liga2">Liga 2</NavDropdown.Item>
                <NavDropdown.Item href="#liga3">Liga 3</NavDropdown.Item>
              </NavDropdown>
              <Nav.Link href="#trf-player">Transfer Pemain</Nav.Link>
              <Nav.Link href="#tim">Tim</Nav.Link>
            </Nav>
            <Form className="d-flex">
              <Form.Control
                type="search"
                placeholder="Search"
                className="me-2"
                aria-label="Search"
              />
              <Button variant="outline-success">Search</Button>
            </Form>
          </Navbar.Collapse>
        </Container>
      </Navbar>
    </>
  );
};

export default NavigationBar;
