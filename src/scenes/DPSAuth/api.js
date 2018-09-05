import { getClient, checkResponse } from 'shared/api';

export async function SetDPSAuthCookie(cookieName, redirectURL) {
  const client = await getClient();
  const response = await client.apis.dps_auth.setDPSAuthCookie({
    cookie_name: cookieName,
    redirect_url: redirectURL,
  });
  checkResponse(response, 'Failed to set DPS Auth cookie');
  return response.body;
}
