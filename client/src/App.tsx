import { useLayoutEffect, useState } from "react";
import { Route, Routes, useLocation } from "react-router-dom";

import NavBar from "components/NavBar";
import Footer from "components/Footer";
import ErrorBoundary from "components/ErrorBoundary";
import Home from "pages/Home";
import ErrorPage from "pages/Error";

import Protected from "components/Auth/Protected";
import Login from "components/Auth/Login";
import Logout from "components/Auth/Logout";

import Schedules from "pages/Schedules";
import Profile from "pages/Profile";
import Users from "pages/Users";

import Container from "@mui/material/Container";

function App() {
  const location = useLocation();
  const [path, setPath] = useState(location.pathname);
  useLayoutEffect(() => setPath(location.pathname), [location.pathname]);
  return (
    <>
      <NavBar />
      <Container maxWidth="xl" component="main" sx={{ mt: 4 }}>
        <ErrorBoundary key={path}>
        <Routes>
          <Route path="/" element={<Home />} />
          <Route path="/login" element={<Login />} />
          <Route path="/logout" element={<Logout />} />

          {/* Protected routes (only for authorized users)  */}
          <Route element={<Protected />}>
            <Route path="/schedules" element={<Schedules />} />
            {/* <Route path="/booking" element={<Booking />} />
                <Route path="/tickets" element={<Tickets />} />
                <Route path="/surveys" element={<Surveys />} /> */}
            <Route path="/profile" element={<Profile />} />
          </Route>
          <Route element={<Protected role="Administrator" />}>
            <Route path="/users" element={<Users />} />
          </Route>

          <Route path="*" element={<ErrorPage code="404" />} />
        </Routes>
        </ErrorBoundary>
      </Container>
      <Footer />
    </>
  );
}

export default App;
