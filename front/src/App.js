import { BrowserRouter as Router, Routes, Route } from "react-router-dom";
import { Home } from "./pages/Home";
import { CircleForm } from "./pages/CircleForm";

function App() {
  return (
    <>
      <Router>
        <Routes>
          <Route path="/" element={<Home />} />
          <Route path="/form" element={<CircleForm />} />
        </Routes>
      </Router>
    </>
  );
}

export default App;
