import axios from "axios";

/**
 * Obtiene nuevo token desde servidor.
 *
 * Retorna promesa que se cumple luego de recibir respuesta desde el servdor.
 * - fulfilled : {string} token
 * - rejected : {object} error
 */
export default async function getToken() {
    let config = {
        method: 'post',
        url: 'http://18.214.103.65:8080/api/auth/login',
        headers: {
            'Content-Type': 'application/json',
            'Accept': 'application/json'
        },
        data: {
            username: "alumnos@alumnos.org",
            password: "m7a/s99"
        }
    }

    try {
        let resp = await axios(config);
        return resp.data.token;
    }
    catch (error) {
        return error.response.data;
    }
}
