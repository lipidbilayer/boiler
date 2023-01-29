import { HandleError } from "../error";

const UserList = ({setData, actions, setLoading}) => {

    if (setLoading != null) {
        setLoading(true)
    }

    fetch(process.env.NEXT_PUBLIC_SERVER_API_URL+'/api/user')
    .then((response) => {
        if(!response.ok){
            throw response
        }
        return response.json()
    })
    .then((data) => {
        data.map((tdata, index) => {
            data[index]['action'] = actions;
        });
        setData(data)
        if (setLoading != null) {
            setLoading(false)
        }
    })
    .catch((error) => {
        if (setLoading != null) {
            setLoading(false)
        }
        HandleError(error)
    });
}

export const UserListTable = ({setData, setLoading}) => {
    
}

export default UserList;