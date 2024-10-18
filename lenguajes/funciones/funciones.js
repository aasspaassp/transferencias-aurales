/* 
Una función nos permite agrupar una secuencia de instrucciones como una unidad que puede ser llamada desde otro 
lugar en un programa, quizás varias veces. Las funciones hacen posible dividir una tarea grande en piezas más pequeñas que
bien podrían ser escritas por diferentes personas separadas tanto por el tiempo como por el espacio.

Una función oculta sus detalles de implementación de sus usuarios. Por todas estas razones, las funciones son 
una parte crítica de cualquier lenguaje de programación.

Una declaración de función tiene un nombre, una lista de parámetros, una lista opcional de resultados y un cuerpo
*/

function devuelveLoMismo(argument) {
    return argument
}

let algo = "energía"

console.log(devuelveLoMismo(algo))

const suenaDo = () => console.log('doooo')
suenaDo()



async function llamar(numero) {
    try {
        console.log("beeep, beeep,beep")
        setTimeout(() => {
            if (numero == 5551111) {
                console.log('holaaaaaaa')
            } else {
                console.log("el número que usted marcó no existe")
            }
        }, 2000)
    } catch (err) {
        console.log(err)
    }
}

llamar(5551111)
llamar(5)
llamar()



