import React from "react";
import Header from "../../Header/Header";

function ProtectedPage(){
    return (
        <>
        <Header/>
        <div>
            <h1>Protected Page</h1>
            <p> You can see this page </p>
        </div>
        </>
    );
}

export default ProtectedPage;