import axios from "axios";

export class Axios {
  async PostService(data: any, sufix: string): Promise<any> {
    const response = await axios.post(`http://localhost:8080/${sufix}`, data);
    return response.data;
  }

  async GetService(data: any, sufix: string): Promise<any> {
    const response = await axios.get(`http://localhost:8080/${sufix}`, {
      params: data,
    });
    return response.data;
  }

  async PutService(data: any, sufix: string): Promise<any> {
    const response = await axios.put(`http://localhost:8080/${sufix}`, data);
    return response.data;
  }
}
