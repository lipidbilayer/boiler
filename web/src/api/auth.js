export const Auth = async (ctx) => {
    let URL = process.env.NEXT_PUBLIC_SERVER_API_URL

    const res = await fetch(URL+'/auth/profile');
    const data = await res.json();
    if
    return data
}

export const Refresh = async(ctx) => {

}