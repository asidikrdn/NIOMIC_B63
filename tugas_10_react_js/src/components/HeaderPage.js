import { Container, Flag, Icon, Input, Image, Header } from "semantic-ui-react";

const HeaderPage = () => {
  return (
    <header>
      <Container style={{ padding: "15px 0px" }}>
        <Flag name="id"></Flag>
        <Icon
          link
          name="angle left"
          size="big"
          style={{ margin: "0px 3em" }}
        ></Icon>
        <Icon link name="angle right" size="big"></Icon>
        <Input
          icon="star"
          placeholder="Search..."
          style={{ width: "50%", margin: "0px 3em" }}
        />
        <Image
          src="https://encrypted-tbn2.gstatic.com/images?q=tbn:ANd9GcQ7xzkB6jECLHRZOS9uRekMHGkkTS4NncEEgzM2QVE8XkrLeffM"
          size="mini"
          circular
          style={{ display: "inline-block" }}
        />
        <Header style={{ display: "inline-block", marginLeft: "1em" }}>
          Patrick
        </Header>
      </Container>
    </header>
  );
};

export default HeaderPage;
