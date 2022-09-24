import { BrowserRouter, Route, Routes } from "react-router-dom";

import NavBar from "components/NavBar";
import Footer from "components/Footer";
import Home from "pages/Home";
import Error from "pages/Error";

import Protected from "components/Auth/Protected";
import Login from "components/Auth/Login";
// import Logout from "components/Auth/Logout";

import Schedules from "pages/Schedules";

import Container from "@mui/material/Container";
import { ThemeProvider } from "@mui/material/styles";
import theme from "utils/theme";

function App() {
  return (
    <>
      <ThemeProvider theme={theme}>
        <BrowserRouter>
          <NavBar />
          <Container maxWidth="xl" component="main">
            <Routes>
              <Route path="/" element={<Home />} />
              <Route path="/login" element={<Login />} />

              {/* Protected routes (only for authorized users)  */}
              <Route element={<Protected />}>
                <Route path="/schedules" element={<Schedules />} />
                {/* <Route path="/booking" element={<Booking />} />
                <Route path="/tickets" element={<Tickets />} />
                <Route path="/surveys" element={<Surveys />} /> */}
                {/* <Route path="/logout" element={<Logout />} /> */}
              </Route>

              <Route path="/403" element={<Error code={403} />} />
              <Route path="*" element={<Error />} />
            </Routes>
          </Container>
          <Footer />
        </BrowserRouter>
      </ThemeProvider>
    </>
  );
}

export default App;
