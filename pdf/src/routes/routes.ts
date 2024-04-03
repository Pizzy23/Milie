import express from "express";
import multer from "multer";
import { PDF_Reader } from "../service/pdf";
import { GPT } from "../service/gpt";
import { PDF_Search } from "../service/search";
import { Axios } from "../service/axios";

export const router = express.Router();
const storage = multer.diskStorage({
  destination: "./pdfsFiles",
  filename: (req, file, cb) => {
    cb(null, file.originalname);
  },
});
const upload = multer({ storage: storage });

const servicePDF = new PDF_Reader();
const serviceGPT = new GPT();
const serviceSearch = new PDF_Search();
const serviceAxios = new Axios();

router.post("/pdf", upload.single("pdf"), async (req, res) => {
  try {
    const uploadedPdfPath = req.file.path;
    const responsible = req.body.responsible;
    const name = req.body.name;
    const number = req.body.number;
    const description = req.body.description;
    const process = req.body.process;

    if (uploadedPdfPath) {
      const t0 = performance.now();
      const array = await servicePDF.extractPages(uploadedPdfPath);
      const t1 = performance.now();
      console.log(`Extraction took ${(t1 - t0 / 1000).toFixed(2)} seconds.`);
      const response = await servicePDF.getText(array);
      res.status(200).json({ message: response });
    } else {
      throw new Error("Pdf Invalid");
    }
  } catch (error) {
    res.status(500).json({ error: "Internal error in pdf." });
  }
});
