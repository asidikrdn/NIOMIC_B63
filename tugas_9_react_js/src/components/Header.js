import {
  Container,
  Dropdown,
  DropdownButton,
  Figure,
  Stack,
} from "react-bootstrap";

const Header = () => {
  return (
    <Container fluid className="px-0 pb-1 mt-2 bg-black">
      <Stack direction="horizontal" gap={2} className="pt-0 align-items-start">
        <DropdownButton id="dropdown-basic-button" title="Pilih Bahasa">
          <Dropdown.Item href="#/action-1">Bahasa Indonesia</Dropdown.Item>
          <Dropdown.Item href="#/action-2">English</Dropdown.Item>
        </DropdownButton>
        <Figure className="ms-auto me-5 mt-4 text-center">
          <Figure.Image
            width={50}
            height={180}
            alt="Alan Saputra"
            src="https://png.pngtree.com/svg/20161113/ef1b24279e.png"
          />
          <Figure.Caption className="text-light fw-semibold">Alan Saputra</Figure.Caption>
        </Figure>
      </Stack>
    </Container>
  );
};

export default Header;
