import { Container, Alert, Breadcrumb } from "react-bootstrap";

const Nav = () => {
  return (
    <Container fluid className="px-0">
      <Alert
        variant="secondary"
        className="ms-auto py-1"
        style={{ width: "35%" }}
      >
        <Breadcrumb className="mt-2">
          <Breadcrumb.Item href="#">Home</Breadcrumb.Item>
          <Breadcrumb.Item href="#">Berita</Breadcrumb.Item>
          <Breadcrumb.Item active>Bola</Breadcrumb.Item>
        </Breadcrumb>
      </Alert>
    </Container>
  );
};

export default Nav;
