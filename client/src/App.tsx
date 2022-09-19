import { BrowserRouter, Route, Routes } from "react-router-dom";
import NavBar from "./components/NavBar";
import Home from "./pages/Home";
import Container from "@mui/material/Container";

function App() {
  return (
    <>
    {/* <Container maxWidth="xl"> */}
      <NavBar />
      <Container maxWidth="xl" component="main">
        Home
      </Container>
    {/* </Container> */}
    </>
    // <Context.Provider value={{ store}}>
    // <BrowserRouter>
    //   <Routes>
    //     <Route path="/" element={<Home />} />
    //   </Routes>
    // </BrowserRouter>
    // </Context.Provider>
  );
}

export default App;
