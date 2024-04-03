import fs from "fs/promises";
import { PDFDocument } from 'pdf-lib'
import pdfParse from "pdf-parse";

export class PDF_Search {
  async extractPageWithContentWord(pdfPath: string, word: string): Promise<{ text: string; originalPage: number }[]> {
    try {
      const documentAsBytes = await fs.readFile(pdfPath);

      const pdfDoc = await PDFDocument.load(documentAsBytes);
      const pagesItContent: { text: string; originalPage: number }[] = [];
      const numberOfPages = pdfDoc.getPageCount();

      for (let i = 0; i < numberOfPages; i++) {
        const subDocument = await PDFDocument.create();

        const copiedPage = await subDocument.copyPages(pdfDoc, [i]);
        subDocument.addPage(copiedPage[0]);

        const pdfBytes = await subDocument.save();
        const text = await this.getPageText(pdfBytes);
        if (text.includes(word)) {
          pagesItContent.push({ text: text, originalPage: i + 1 });
        }
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
}
