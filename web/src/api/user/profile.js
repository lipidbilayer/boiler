export const UserProfileAPI = async () => {
    let URL = process.env.NEXT_PUBLIC_SERVER_API_URL

    const res = await fetch(URL+'/auth/profile');
    const data = await res.json();
    return data
}

const UserProfile = ({setData, setLoading}) => {
    if (setLoading !== undefined) {
        setLoading(true)
    }


    fetch(process.env.NEXT_PUBLIC_SERVER_API_URL+'/auth/profile')
    .then((response) => response.json())
    .then((data) => {
        setData(data)
        if (setLoading !== undefined) {
            setLoading(false)
        }
    })
    .catch((error) => {
        if (setLoading !== undefined) {
            setLoading(false)
        }
    });
}

export default UserProfile;