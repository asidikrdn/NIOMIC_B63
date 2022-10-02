import { Container, Alert, Button, Stack, Badge } from "react-bootstrap";

const Header = () => {
  return (
    <Container fluid className="px-0 py-3">
      <Stack
        direction="horizontal"
        gap={2}
        className="d-flex justify-content-center"
      >
        <Alert variant="primary" className="w-25 text-dark my-0">
          Website React Bootstrap
        </Alert>

        <Button className="ms-auto">
          Notification
          <Badge bg="light" text="dark" className="ms-2">
            9
          </Badge>
        </Button>
        <Button className="me-2">
          Message
          <Badge bg="light" text="dark" className="ms-2">
            19
          </Badge>
        </Button>

        <img
          src="https://banner2.kisspng.com/20180420/kuq/kisspng-computer-icons-user-clip-art-my-vector-5ad9a348cd2ce1.8035052415242125528404.jpg"
          alt="User"
          style={{ width: "50px" }}
          className="mx-2"
        />
        <h1 className="me-4">Andreas</h1>
      </Stack>
    </Container>
  );
};

export default Header;
