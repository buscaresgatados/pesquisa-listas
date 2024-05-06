// services/ApiService.js
import axios from 'axios';

const ResultService = {
  async search(termo) {
    try {
      const response = await axios.get(`https://refugio-rs-prd-d5zml3w7fa-rj.a.run.app/pessoa?nome=${termo}`);
      return response.data;
    } catch (error) {
      throw new Error(`Erro ao buscar dados: ${error}`);
    }
  }
};

export default ResultService;
