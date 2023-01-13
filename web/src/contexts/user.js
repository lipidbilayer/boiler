import { createContext, useState, useEffect } from 'react';
import UserProfile from '../api/user/profile';

const UserContext = createContext(0);

export const UserProvider = ({ children }) => {
  const [user, setUser] = useState({});


  useEffect(() => {
    UserProfile(setUser, null)
    }, [])


  return (
    <UserContext.Provider value={[user]}>
      {children}
    </UserContext.Provider>
  );
};

export default UserContext;