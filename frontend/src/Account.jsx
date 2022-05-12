import { useEffect, useState } from "react";
import axios from "axios"

const Account = (props) => {
	const [accountId, setAccountId] = useState(1);
	const [user, setUser] = useState({});

	useEffect( () => {
		console.log("Call effect");
		axios.get('/api/accounts/' + accountId)
		.then(response => {
			setUser(response.data)
		}).catch(function (error) {
			setUser({ID:-1, name: 'Account not found' })
		});
	}, [accountId]);

	
	return( 
		<div>
			Account ID : <input type="text" onChange={(e)=> setAccountId(e.target.value) }/>
			<p>Account ID = {user.ID}</p>
			<p>Account Name = {user.name}</p>
			<h1>Hello Workshop with testing #1</h1>
		</div>
	);
}

export default Account;