import React, { useState } from "react";
import {
  Container,
  Alert,
  Breadcrumb,
  Button,
  Collapse,
  Tabs,
  Tab,
  Table,
  Spinner,
  Stack,
  Pagination,
  ProgressBar,
  Popover,
  OverlayTrigger,
} from "react-bootstrap";

const popover = (
  <Popover id="popover-basic">
    <Popover.Header as="h3">Informasi Website</Popover.Header>
    <Popover.Body>
      Website ini dibuat untk memudahkan dalam pencarian informasi terkini
      tentang berita olahraga
    </Popover.Body>
  </Popover>
);

const MainContent = () => {
  const [showVersiWebsite, setShowVersiWebsite] = useState(false);

  return (
    <Container className="pb-5 pt-0 px-0">
      {/* Breadcumb */}
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
      {/* End of Breadcumb */}

      <Stack direction="horizontal" className="justify-content-between">
        <h2 className="d-inline-block">Rumor Transfer Pemain</h2>
        <div>
          <Spinner animation="border" variant="success" />
          <Spinner animation="grow" variant="success" />
        </div>
      </Stack>
      <Tabs
        defaultActiveKey="profile"
        id="fill-tab-example"
        className="mb-3"
        fill
      >
        <Tab eventKey="home" title="Semua Transfer">
          {/* <Sonnet /> */}
          <Table hover>
            <thead>
              <tr>
                <th>#</th>
                <th>Nama Pemain</th>
                <th>Tim</th>
                <th>Transfer</th>
                <th>Proses Transfer</th>
              </tr>
            </thead>
            <tbody>
              <tr>
                <td>1</td>
                <td>MAROUANE FELLAINI</td>
                <td>MANCHESTER UNITED</td>
                <td>SHANDONG LUNENG</td>
                <td>
                  <ProgressBar animated now={85} label={`85%`} />
                </td>
              </tr>
              <tr>
                <td>2</td>
                <td>LUIS NANI</td>
                <td>SPORTING CP</td>
                <td>ORLANDO CITY</td>
                <td>
                  <ProgressBar animated now={55} label={`55%`} />
                </td>
              </tr>
              <tr>
                <td>3</td>
                <td>MAREK HAMSIK</td>
                <td>NAPOLI</td>
                <td>DALIAN YIFANG</td>
                <td>
                  <ProgressBar animated now={95} label={`95%`} />
                </td>
              </tr>
              <tr>
                <td>4</td>
                <td>SARDAR AZMOUN</td>
                <td>RUBIN KAZAN</td>
                <td>ZENIT ST PETERSBURG</td>
                <td>
                  <ProgressBar animated now={100} label={`100%`} />
                </td>
              </tr>
              <tr>
                <td>5</td>
                <td>MICHY BATSHUAYI</td>
                <td>CHELSEA</td>
                <td>CRYSTAL PALACE</td>
                <td>
                  <ProgressBar animated now={50} label={`50%`} />
                </td>
              </tr>
              <tr>
                <td>6</td>
                <td>LUCAS PIAZON</td>
                <td>CHELSEA</td>
                <td>CHIEVO</td>
                <td>
                  <ProgressBar animated now={100} label={`100%`} />
                </td>
              </tr>
            </tbody>
          </Table>
          <Pagination>
            <Pagination.First />
            <Pagination.Prev />
            <Pagination.Item active>{1}</Pagination.Item>
            <Pagination.Ellipsis />

            <Pagination.Item>{10}</Pagination.Item>
            <Pagination.Item>{11}</Pagination.Item>
            <Pagination.Item>{12}</Pagination.Item>
            <Pagination.Item>{13}</Pagination.Item>
            <Pagination.Item disabled>{14}</Pagination.Item>

            <Pagination.Ellipsis />
            <Pagination.Item>{20}</Pagination.Item>
            <Pagination.Next />
            <Pagination.Last />
          </Pagination>
        </Tab>
        <Tab eventKey="1" title="Liga Primer Inggris">
          {/* <Sonnet /> */}
        </Tab>
        <Tab eventKey="2" title="Serie A">
          {/* <Sonnet /> */}
        </Tab>
        <Tab eventKey="3" title="Divisi Primera">
          {/* <Sonnet /> */}
        </Tab>
        <Tab eventKey="4" title="Bundesliga">
          {/* <Sonnet /> */}
        </Tab>
        <Tab eventKey="5" title="Liga 1 Indonesia">
          {/* <Sonnet /> */}
        </Tab>
      </Tabs>
      <OverlayTrigger trigger="click" placement="right" overlay={popover}>
        <Button className="me-2">Informasi</Button>
      </OverlayTrigger>
      {/* <div className="d-inline-block"> */}
      <Button onClick={() => setShowVersiWebsite(!showVersiWebsite)}>
        Versi Website
      </Button>
      <Collapse in={showVersiWebsite} style={{marginLeft:'100px'}}>
        <div id="example-collapse-text">Akses Sport V1.0</div>
      </Collapse>
      {/* </div> */}
    </Container>
  );
};

export default MainContent;
