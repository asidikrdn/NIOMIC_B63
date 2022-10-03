import {
  List,
  Loader,
  Dimmer,
  Placeholder,
  Segment,
  Grid,
  Divider,
  Header,
  Icon,
  Search,
  Button,
  Image,
} from "semantic-ui-react";

const App = () => {
  return (
    <>
      <Segment placeholder style={{ marginTop: "3em" }}>
        <Grid columns={2} stackable textAlign="center">
          <Divider vertical>Or</Divider>

          <Grid.Row verticalAlign="middle">
            <Grid.Column>
              <Header icon>
                <Icon name="search" />
                Cari document
              </Header>

              <Search placeholder="Search document..." />
            </Grid.Column>

            <Grid.Column>
              <Header icon>
                <Icon name="file pdf outline" />
                Tambah Document Baru
              </Header>
              <Button primary>Create Document</Button>
            </Grid.Column>
          </Grid.Row>
        </Grid>
      </Segment>

      <Segment>
        <Dimmer active>
          <Loader content="Loading" />
        </Dimmer>

        <Image src="https://react.semantic-ui.com/images/wireframe/short-paragraph.png" />
      </Segment>

      <Grid columns={3} stackable>
        <Grid.Column>
          <Segment raised>
            <Placeholder>
              <Placeholder.Header image>
                <Placeholder.Line />
                <Placeholder.Line />
              </Placeholder.Header>
              <Placeholder.Paragraph>
                <Placeholder.Line length="medium" />
                <Placeholder.Line length="short" />
              </Placeholder.Paragraph>
            </Placeholder>
          </Segment>
        </Grid.Column>

        <Grid.Column width={7}>
          <Segment raised>
            <Placeholder>
              <Placeholder.Header image>
                <Placeholder.Line />
                <Placeholder.Line />
              </Placeholder.Header>
              <Placeholder.Paragraph>
                <Placeholder.Line length="medium" />
                <Placeholder.Line length="short" />
              </Placeholder.Paragraph>
            </Placeholder>
          </Segment>
        </Grid.Column>

        <Grid.Column width={2}>
          {/* <Segment raised> */}
          <Header>Website Terkait</Header>
          <List>
            <List.Item>
              <List.Icon name="linkify" />
              <List.Content>
                <a href="http://www.google.co.id">Google</a>
              </List.Content>
            </List.Item>
            <List.Item>
              <List.Icon name="linkify" />
              <List.Content>
                <a href="http://www.facebook.com">Facebook</a>
              </List.Content>
            </List.Item>
            <List.Item>
              <List.Icon name="linkify" />
              <List.Content>
                <a href="http://www.semantic-ui.com">Semantic UI</a>
              </List.Content>
            </List.Item>
            <List.Item>
              <List.Icon name="linkify" />
              <List.Content>
                <a href="http://www.niomic.id">Niomic</a>
              </List.Content>
            </List.Item>
            <List.Item>
              <List.Icon name="linkify" />
              <List.Content>
                <a href="http://www.reactjs.org">React</a>
              </List.Content>
            </List.Item>
          </List>
          {/* </Segment> */}
        </Grid.Column>
      </Grid>
    </>
  );
};

export default App;
