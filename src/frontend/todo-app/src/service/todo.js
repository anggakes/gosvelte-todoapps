// Api.js
import axios from "axios";

// Create a instance of axios to use the same base url.
const axiosAPI = axios.create({
    baseURL : "http://localhost:8181/todos/" // it's not recommended to have this info here.
});

// implement a method to execute all the request from here.
const apiRequest = (method, url, request) => {
    const headers = {
        authorization: ""
    };
    //using the axios instance to perform the request that received from each http method
    return axiosAPI({
        method,
        url,
        data: request,
        headers
    }).then(res => {
        return Promise.resolve(res.data);
    })
        .catch(err => {
            return Promise.reject(err);
        });
};

// function to execute the http get request
const get = (url, request) => apiRequest("get",url,request);

// function to execute the http delete request
const deleteRequest = (url, request) =>  apiRequest("delete", url, request);

// function to execute the http post request
const post = (url, request) => apiRequest("post", url, request);

// function to execute the http put request
const put = (url, request) => apiRequest("put", url, request);

// function to execute the http path request
const patch = (url, request) =>  apiRequest("patch", url, request);

// expose your method to other services or actions
const TodoService ={
    list : () =>{
        return axios.get('http://localhost:8181/todos/?size=3')
            .then(res => {
                return Promise.resolve(res.data);
            })
            .catch(err => {
                return Promise.reject(err);
            });
    },
    create: (todo) =>  {
        return axios.post('http://localhost:8181/todos/', todo)
            .then(res => {
                return Promise.resolve(res.data);
            })
            .catch(err => {
                return Promise.reject(err);
            });
    },
    update: (todo) => {
        return axios.put('http://localhost:8181/todos/'+todo.id+"/", todo)
            .then(res => {
                return Promise.resolve(res.data);
            })
            .catch(err => {
                return Promise.reject(err);
            });
    },
    delete: (id) => {
        return axios.delete('http://localhost:8181/todos/'+id+"/")
            .then(res => {
                return Promise.resolve(res.data);
            })
            .catch(err => {
                return Promise.reject(err);
            });
    }
};
export default TodoService;