class SatelliteError extends Error {};

export async function register({ fullName, shortName, email, password }) {
    const response = await fetch("/api/v0/auth/register", {
        method: "POST",
        headers: {
            "Content-Type": "application/json"
        },
        body: JSON.stringify({
            fullName,
            shortName,
            email,
            password,
            isProfessional: false,
            position: "dj",
            secret: ""
        })
    });

    const token = await response.json();

    if(typeof token !== "string") {
        throw new SatelliteError(`Satellite register returned bad value ${JSON.stringify(token)}`)
    }

    return token;
}