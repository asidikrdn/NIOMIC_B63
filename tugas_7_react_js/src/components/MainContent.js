import { Container, Card, Button, Stack } from "react-bootstrap";

const MainContent = () => {
  return (
    <Container className="py-5">
      <Stack direction="horizontal" className="align-items-start" gap={3}>
        <Card style={{ width: "18rem" }}>
          <Card.Img
            variant="top"
            src="https://images.performgroup.com/di/library/GOAL/f7/ad/diego-godin-vs-thailand_l9viaa2o7afh19gjl6vnekbgq.jpg?t=-1214473334&w=1920&h=1362&quality=100"
          />
          <Card.Body>
            <Card.Title>
              Diego Goldin Ukir Rekor Spesial Lawan Thailand
            </Card.Title>
            <Card.Text>
              Godin sekarang menjadi pemain Urugiay dengan caps terbanyak
              setelah membantu negaranya melibas Thailand di final Piala
              Tiongkok. Diego Goldin boleh berbangga setelah resmi menjadi...
            </Card.Text>
            <Button variant="primary">Read More</Button>
          </Card.Body>
        </Card>
        <Card style={{ width: "18rem" }}>
          <Card.Img
            variant="top"
            src="https://images.performgroup.com/di/library/GOAL/30/f0/neymar-psg-paris-saint-germain-2018-19_gmakk5891n9c19udcwofmrsgc.jpg?t=-1258742791&quality=60&w=1600"
          />
          <Card.Body>
            <Card.Title>
              Neymar Bahas Kontrak Baru Di Paris Saint-Germaint
            </Card.Title>
            <Card.Text>
              Ayah sang pemain bintang Brasil berusaha meredam rumor
              ketertarikan Madrid. Neymar sedang berdiskusi dengan Paris
              Saint-Germain soal perpanjangan kontrak bersama kampiun League 1
              Prancis, demikian kata ayahnya...
            </Card.Text>
            <Button variant="primary">Read More</Button>
          </Card.Body>
        </Card>
        <Card style={{ width: "18rem" }}>
          <Card.Img
            variant="top"
            src="https://images.performgroup.com/di/library/GOAL/f3/33/england-celebrate-vs-montenegro_e97kv5anpqlq1gr9nc2qaz4ko.jpg?t=-1232521935&quality=60&w=1600"
          />
          <Card.Body>
            <Card.Title>Inggris Mau Jadi Tim Terbaik Di Dunia</Card.Title>
            <Card.Text>
              Kumpulan talenta terbaik yang dimiliki Inggris saat ini membuat
              Barkley sangat optimistis. Ross Barkley merasa Inggris sudah
              berada di jalur yang benar untuk bisa menjadi tim terbaik di dunia
              dengan banyaknya talenta yang muncul sekarang ini...
            </Card.Text>
            <Button variant="primary">Read More</Button>
          </Card.Body>
        </Card>
        <Card style={{ width: "18rem" }}>
          <Card.Img
            variant="top"
            src="https://images.performgroup.com/di/library/GOAL/63/7f/sani-rizki-fauzi-indonesia-u-22_4yztc4nehkt61fa7omuyagf71.jpg?t=351764441&quality=60&w=1600"
          />
          <Card.Body>
            <Card.Title>
              Sani Rizki Fauzi Tegaskan Mental Pemain Timnas Indonesia U-23
              Tetap Bagus
            </Card.Title>
            <Card.Text>
              Gelandang tim nasional (timnas) Indonesia U-23, Sani Rizki Fauzi,
              menegaskan kesiapannya untuk bermain pada laga melawan Brunei
              Darussalam U-23. Kedua tim akan bentrok pada laga pamungkas grup
              K...
            </Card.Text>
            <Button variant="primary">Read More</Button>
          </Card.Body>
        </Card>
      </Stack>
    </Container>
  );
};

export default MainContent;
