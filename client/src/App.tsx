import { BrowserRouter, Route, Routes } from "react-router-dom";
import NavBar from "./components/NavBar";
import Home from "./pages/Home";
import NotFound from "./pages/NotFound";
import Schedules from "./pages/Schedules";
import Container from "@mui/material/Container";

function App() {
  return (
    <>
      <BrowserRouter>
        <NavBar />
        <Container maxWidth="xl" component="main">
          <Routes>
            <Route path="/" element={<Home />} />
            <Route path="/schedules" element={<Schedules />} />
            {/* <Route path="/booking" element={<Booking />} />
                <Route path="/tickets" element={<Tickets />} />
                <Route path="/surveys" element={<Surveys />} /> */}
            <Route path="*" element={<NotFound />} />
          </Routes>
        </Container>
      </BrowserRouter>
    </>
  );
}

export default App;
