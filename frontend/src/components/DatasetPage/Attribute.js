import React from "react";

const Attribute = ({ attributeName, value }) => {
	return (
		<div className="w-full flex">
			<p className="text-gray-800 mr-4">{attributeName}: </p>
			<p className="font-medium">{value}</p>
		</div>
	);
};

export default Attribute;
