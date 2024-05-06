// services/ApiService.js
import axios from 'axios';



const ResultService = {  
  async search(termo) {
    try {
      const response = await axios.get(`https://api.exemplo.com/busca?termo=${termo}`);
      return response.data;
    } catch (error) {
      throw new Error(`Erro ao buscar dados: ${error}`);
    }
  }
};

export default ResultService;
