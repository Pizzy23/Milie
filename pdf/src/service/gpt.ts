import axios from "axios";
import "dotenv/config";

const key = process.env.GPT;

export class GPT {
  async gptResume(message: string): Promise<string[]> {
    try {
      const response = await axios.post<{
        choices: { message: { content: string } }[];
      }>(
        "https://api.openai.com/v1/chat/completions",
        {
          model: "gpt-3.5-turbo",
          messages: [
            {
              role: "system",
              content: "Você é um assistente extremamente útil.",
            },
            {
              role: "user",
              content: `APENAS RESUME E NÃO ENSINE A RESUMAR, Faça-me um resumo em 10 partes de 50 palavras: ${message}`,
            },
          ],
        },
        {
          headers: {
            "Content-Type": "application/json",
            Authorization: `Bearer ${key}`,
          },
        }
      );

      const res = response.data.choices[0].message.content;
      const cleanedText = res.trim().replace(/^"|"$/g, "");
      const clearNumber = cleanedText.replace(/\d+\./g, "");
      const lines = clearNumber.split("\n");

      return lines;
    } catch (error) {
      throw new Error(`Error sending chat request: ${error}`);
    }
  }

  async gptDiagram(
    message: string
  ): Promise<{ itemId: number; ItemName: string; blockedBy: number }[]> {
    try {
      const response = await axios.post<{
        choices: { message: { content: string } }[];
      }>(
        "https://api.openai.com/v1/chat/completions",
        {
          model: "gpt-3.5-turbo",
          messages: [
            { role: "system", content: "You are a helpful assistant." },
            {
              role: "user",
              content: `Summarize this in several steps 1, 2, 3, 4... and so on, as you would explain in a flowchart ${message} summary (max 20 steps)`,
            },
          ],
        },
        {
          headers: {
            "Content-Type": "application/json",
            Authorization: `Bearer ${key}`,
          },
        }
      );
      const res = response.data.choices[0].message.content;
      const cleanedText = res.trim().replace(/^"|"$/g, "");
      const clearNumber = cleanedText.replace(/\d+\./g, "");
      const lines = clearNumber.split("\n");
      const arrayObj: {
        itemId: number;
        ItemName: string;
        blockedBy: number;
      }[] = [];
      for (let i = 0; i <= lines.length - 1; i++) {
        arrayObj.push({
          itemId: i + 1,
          ItemName: lines[i],
          blockedBy: i,
        });
      }
      return arrayObj;
    } catch (error) {
      throw new Error(`Error sending chat request: ${error}`);
    }
  }
}