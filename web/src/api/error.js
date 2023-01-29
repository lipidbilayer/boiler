
import Router from 'next/router'

export const HandleError = (error) => {
    console.log(error)
    if(error instanceof Response){
        if(error.status == 401){
            Router.push('/login')
        }
    }
}