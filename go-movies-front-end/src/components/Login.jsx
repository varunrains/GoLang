import React, { useState } from "react";
import { useNavigate, useOutletContext } from "react-router-dom";
import Input from "./form/Input";

const Login = () => {
  const [email, setEmail] = useState("");
  const [password, setPassword] = useState("");

  const { setJwtToken } = useOutletContext();
  const { setAlertClassName } = useOutletContext();
  const { setAlertMessage } = useOutletContext();
  const {toggleRefresh} = useOutletContext();

  const navigate  = useNavigate();

  const handleSubmit = async (event) => {
    event.preventDefault();

    //build the request payload
    let payload = {
      email: email,
      password: password,
    }

    const requestOptions = {
      method:"POST",
      headers:{
        'Content-Type':'application/json'
      },
      credentials: 'include',
      body:JSON.stringify(payload),
    }

    try {
      var response = await fetch("http://localhost:8080/authenticate", requestOptions);
      var data = await response.json();

      if(data?.message === 'invalid credentials'){
        setAlertClassName("alert-danger");
        setAlertMessage(data?.message);
        return;
      }
toggleRefresh(true);
      setJwtToken(data.access_token);
      setAlertClassName("d-none");
      setAlertMessage("");
      navigate("/");
    } catch (err) {
      setAlertClassName("alert-danger");
      setAlertMessage(err);
      console.log(err);
    }
  };

  return (
    <div className="col-md-6 offset-md-3">
      <h2>Login</h2>
      <hr />

      <form onSubmit={handleSubmit}>
        <Input
          title="Email Address"
          type="email"
          className="form-control"
          name="email"
          autoComplete="email-new"
          onChange={(event) => setEmail(event.target.value)}
        />
        <Input
          title="Password"
          type="password"
          className="form-control"
          name="password"
          autoComplete="password-new"
          onChange={(event) => setPassword(event.target.value)}
        />
        <input type="submit" className="btn btn-primary" value="Login"></input>
      </form>
    </div>
  );
};

export default Login;
