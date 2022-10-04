import {
  Slider,
  Slide,
  Caption,
  Row,
  Col,
  Card,
  Icon,
  CardTitle,
} from "react-materialize";

const App = () => {
  return (
    <>
      <Slider
        fullscreen={false}
        options={{
          duration: 500,
          height: 400,
          indicators: true,
          interval: 6000,
        }}
      >
        <Slide
          image={
            <img
              alt=""
              src="https://s-ec.bstatic.com/images/hotel/max1024x768/106/106114417.jpg"
            />
          }
        >
          <Caption placement="center" style={{ textShadow: "0 0 5px black" }}>
            <h3>Pantai</h3>
            <h5 className="light grey-text text-lighten-3">
              Dapatkan Kesempatan Jalan - Jalan Ke Pantai Dengan Membeli Sebuah
              Product
            </h5>
          </Caption>
        </Slide>
        <Slide
          image={
            <img
              alt=""
              src="https://cdn.antaranews.com/cache/800x533/2020/10/19/hong-kong-1081704_960_720.jpg"
            />
          }
        >
          <Caption placement="center" style={{ textShadow: "0 0 5px black" }}>
            <h3>Hongkong</h3>
            <h5 className="light grey-text text-lighten-3">
              Dapatkan Kesempatan Jalan - Jalan Ke Hongkong Dengan Membeli
              Sebuah Product
            </h5>
          </Caption>
        </Slide>
        <Slide
          image={
            <img
              alt=""
              src="https://photos.mandarinoriental.com/is/image/MandarinOriental/paris-2017-home?wid=2880&hei=1280&fmt=jpeg&crop=9,336,2699,1200&anchor=1358,936&qlt=75,0&fit=wrap&op_sharpen=0&resMode=sharp2&op_usm=0,0,0,0&iccEmbed=0&printRes=72"
            />
          }
        >
          <Caption placement="center" style={{ textShadow: "0 0 5px black" }}>
            <h3>Paris</h3>
            <h5 className="light grey-text text-lighten-3">
              Dapatkan Kesempatan Jalan - Jalan Ke Paris Dengan Membeli Sebuah
              Product
            </h5>
          </Caption>
        </Slide>
      </Slider>

      <h5>Hot List</h5>
      <Row>
        <Col m={3} s={6}>
          <Card
            className="small"
            closeIcon={<Icon>close</Icon>}
            header={
              <CardTitle image="https://ecs7.tokopedia.net/img/cache/300/attachment/2018/8/15/25749302/25749302_274a89a6-ea41-4a3a-ac6b-ce2881016369.jpg">
                <img src="" alt="" />
                Audio
              </CardTitle>
            }
            revealIcon={<Icon>more_vert</Icon>}
          >
            Audio Technica - Mulai Dari <span className="red">Rp 910 rb</span>
          </Card>
        </Col>
        <Col m={3} s={6}>
          <Card
            className="small"
            closeIcon={<Icon>close</Icon>}
            header={
              <CardTitle image="https://ecs7.tokopedia.net/img/cache/300/attachment/2019/1/16/25287362/25287362_a79f0aa6-cf1f-484f-b1a8-83993245cd6f.jpg">
                Handphone
              </CardTitle>
            }
            revealIcon={<Icon>more_vert</Icon>}
          >
            Handphone Samsung - Mulai Dari{" "}
            <span className="red">Rp 200 rb</span>
          </Card>
        </Col>
        <Col m={3} s={6}>
          <Card
            className="small"
            closeIcon={<Icon>close</Icon>}
            header={
              <CardTitle image="https://ecs7.tokopedia.net/img/cache/300/attachment/2018/6/26/28201197/28201197_856195f2-b0ba-4711-8200-fd95beaf1df5.jpg">
                Sequishy
              </CardTitle>
            }
            revealIcon={<Icon>more_vert</Icon>}
          >
            Sequishy - Mulai Dari <span className="red">Rp 5,5 rb</span>
          </Card>
        </Col>
        <Col m={3} s={6}>
          <Card
            className="small"
            closeIcon={<Icon>close</Icon>}
            header={
              <CardTitle image="https://ecs7.tokopedia.net/img/cache/300/attachment/2018/4/23/28201197/28201197_0383072a-1766-453c-951a-8b53a719cff0.jpg">
                Koleksi The Avenger
              </CardTitle>
            }
            revealIcon={<Icon>more_vert</Icon>}
          >
            Koleksi The Avenger - Mulai Dari{" "}
            <span className="red">Rp 10 rb</span>
          </Card>
        </Col>
      </Row>
    </>
  );
};

export default App;
