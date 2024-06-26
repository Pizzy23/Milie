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

    if (uploadedPdfPath) {
      const t0 = performance.now();
      const array = await servicePDF.extractPages(uploadedPdfPath);
      const t1 = performance.now();
      console.log(`Extraction took ${(t1 - t0 / 1000).toFixed(2)} seconds.`);
      res.status(200).json({ message: array });
    } else {
      throw new Error("Pdf Invalid");
    }
  } catch (error) {
    res.status(500).json({ error: "Internal error in pdf." });
  }
});

router.post("/text-extract", async (req, res) => {
  try {
    const text = await servicePDF.getText();
    const res = await serviceAxios.PostService(text, "add-questions");
    res.status(200).json(text);
  } catch (e) {
    res.status(500).json({ error: "Internal error in parse texts" });
  }
});
