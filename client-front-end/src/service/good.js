
import axios from '../utils/axios'

export function getDetail(id) {
  return axios.get(`/goods/detail/${id}`);
}
export function getReview(id,params) {
  return axios.get(`/goods/review/${id}`,{params});
}

export function getCategory() {
  return axios.get('/categories');
}

export function search(params) {
  return axios.get('/search', { params });
}

