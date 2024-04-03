import express from "express";
import bodyParser from "body-parser";
import { router } from "./routes/routes";
import morgan from "morgan";
import cors from "cors";

const app = express();
const port = process.env.PORT || 3000;

app.use(bodyParser.json());
app.use(morgan("dev"));
app.use(cors());

app.use("/", router);

app.listen(port, () => {
  console.log(`Server run : ${port}`);
});
