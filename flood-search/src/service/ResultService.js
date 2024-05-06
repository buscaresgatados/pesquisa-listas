import axios from 'axios';
const apiUrl = "https://refugio-rs-prd-d5zml3w7fa-rj.a.run.app/"

const ResultService = {  
  async search(termo) {
    try {
      const response = await axios.get(`${apiUrl}/pessoa?nome=${termo}`);
      return response.data;
    } catch (error) {
      throw new Error(`Erro ao buscar dados: ${error}`);
    }
  }
};

export default ResultService;
