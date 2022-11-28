import "./App.css";
import axios from "axios";
import { useEffect } from "react";

//(☞ ͡° ͜ʖ ͡°)☞

function App() {
  const getProducts = async () => {
    await axios
      .get("http://localhost:4000/products")
      .then((p) => console.log(p))
      .catch((err) => console.log(err));
  };

  useEffect(() => {
    getProducts();
  }, []);

  return (
    <div className="App">
      <h1>Hello World</h1>
    </div>
  );
}

export default App;
