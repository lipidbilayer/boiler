import { createContext, useState, useEffect } from 'react';
import { UserProfileAPI } from '../api/user/user';

const UserContext = createContext(0, () => {});

export const UserProvider = ({ children }) => {
  const [user, setUser] = useState({});


  useEffect(() => {
    let data = UserProfileAPI()
    setUser(data);
    }, [])


  return (
    <UserContext.Provider value={[user]}>
      {children}
    </UserContext.Provider>
  );
};

export default UserContext;