import axios from "axios";

/**
 * Obtiene los datos del sensor Radiometro.
 *
 * Parametros:
 * - {string} tkn : token de autenticacion
 *
 * Retorna promesa que se cumple luego de recibir respuesta desde el servdor.
 * - fulfilled : {array} arreglo de datos
 * - rejected : {object} error
 *
 * Los elementos del arreglo son objetos con la siguiente estructura:
 * ```
 * {
 *   ts: {int},
 *   date: {string},
 *   hour: {string},
 *   value: {string}
 * }
 * ```
 */
export default async function getRadiometro(tkn) {
    let config = {
        method: 'get',
        url: 'http://18.214.103.65:8080/api/plugins/telemetry/DEVICE/efdd9590-4550-11ed-b4b1-1bcb8f5daa77/values/timeseries?keys=Fecha,Hora,Albedo&startTs=1265046352083&endTs=1665044947549',
        headers: {
            'Content-Type': 'application/json',
            'X-Authorization': 'Bearer ' + tkn
        }
    };

    try {
        let resp = await axios(config);

        let fechas = resp.data.Fecha;
        let horas  = resp.data.Hora;
        let datos  = resp.data.Albedo;

        for(let i = 0; i < datos.length; i++){
            if(datos[i].value === '/'){
                datos[i].value = '-1';
            }
            datos[i].date = fechas[i].value;
            datos[i].hour = horas[i].value;
        }

        return datos;
    }
    catch (error) {
        return error.response.data;
    }
}
