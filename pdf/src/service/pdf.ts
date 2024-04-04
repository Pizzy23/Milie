import fs from "fs/promises";
import { PDFDocument } from "pdf-lib";
import pdfParse from "pdf-parse";
import { text1 } from "./text";

export class PDF_Reader {
  async extractPages(pdfPath: string): Promise<string[]> {
    try {
      const documentAsBytes = await fs.readFile(pdfPath);

      const pdfDoc = await PDFDocument.load(documentAsBytes);
      const pagesItContent: string[] = [];
      const numberOfPages = pdfDoc.getPageCount();

      for (let i = 0; i < numberOfPages; i++) {
        const subDocument = await PDFDocument.create();

        const copiedPage = await subDocument.copyPages(pdfDoc, [i]);
        subDocument.addPage(copiedPage[0]);

        const pdfBytes = await subDocument.save();
        const text = await this.getPageText(pdfBytes);

        pagesItContent.push(text);
      }
      return pagesItContent;
    } catch (error) {
      throw new Error(`Error splitting PDF: ${error.message}`);
    }
  }

  async getPageText(pdfBytes: Uint8Array): Promise<string> {
    const pdfData = await pdfParse(pdfBytes);
    const pdfText = pdfData.text;
    const pdfClear = pdfText.replace(/\n/g, " ");
    return pdfClear;
  }

  async arrayForText(array: string[]): Promise<string> {
    try {
      let textNew = "";
      for (let i = 0; i <= array.length - 1; i++) {
        textNew += array[i];
      }
      return textNew;
    } catch (error) {
      throw error;
    }
  }
  async getText(): Promise<any> {
    const texts = [text1];
    const categorys = {};

    texts.forEach(function (text) {
      const linhas = text.trim().split("\n");

      let categoryNow = "";

      for (let i = 0; i < linhas.length; i++) {
        const linha = linhas[i].trim();

        if (linha.includes("–")) {
          categoryNow = linha.trim();
          categorys[categoryNow] = [];
        } else if (linha.match(/^\d+\s/)) {
          const partes = linha.split(" ");
          const number = partes.shift();
          const desc = partes.join(" ");
          categorys[categoryNow].push({ number, desc });
        }
      }
    });

    return categorys;
  }
}
