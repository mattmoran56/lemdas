import { useState, useEffect } from "react";
import { jwtDecode } from "jwt-decode";

const useAuth = () => {
	const [loading, setLoading] = useState(true);
	const [error, setError] = useState(null);
	const [username, setUsername] = useState(null);

	useEffect(() => {
		const token = localStorage.getItem("token");

		if (token) {
			try {
				const decodedToken = jwtDecode(token);
				const { Name } = decodedToken;
				setUsername(Name);
			} catch (err) {
				setError("Invalid token");
			}
		}

		setLoading(false);
	}, []);

	return { username, error, loading };
};

export default useAuth;
