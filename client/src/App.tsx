import { BrowserRouter, Route, Routes } from "react-router-dom";

import NavBar from "components/NavBar";
import Home from "pages/Home";
import Error from "pages/Error";

import Protected from "components/Auth/Protected";
import Login from "components/Auth/Login";
// import Logout from "components/Auth/Logout";

import Schedules from "pages/Schedules";

import Container from "@mui/material/Container";

function App() {
  return (
    <>
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

            <Route path="*" element={<Error />} />
          </Routes>
        </Container>
      </BrowserRouter>
    </>
  );
}

export default App;
