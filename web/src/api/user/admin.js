const UserList = ({setData, actions, setLoading}) => {
    if (setLoading != null) {
        setLoading(true)
    }


    fetch(process.env.NEXT_PUBLIC_SERVER_API_URL+'/api/user')
    .then((response) => {
        console.log(response)
        response.json()
    })
    .then((data) => {
        data.map((tdata, index) => {
            data[index]['action'] = actions;
        });
        setData(data)
        if (setLoading != nulld) {
            setLoading(false)
        }
    })
    .catch((error) => {
        // console.log(error);
        if (setLoading != null) {
            setLoading(false)
        }
    });
}

export const UserListTable = ({setData, setLoading}) => {
    
}

export default UserList;