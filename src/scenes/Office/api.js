import { getClient, checkResponse } from 'shared/api';

// MOVE QUEUE
export async function RetrieveMovesForOffice(queueType) {
  const client = await getClient();
  const response = await client.apis.queues.showQueue({
    queueType,
  });
  checkResponse(response, 'failed to retrieve moves due to server error');
  return response.body;
}

// MOVE
export async function LoadMove(moveId) {
  const client = await getClient();
  const response = await client.apis.moves.showMove({
    moveId,
  });
  checkResponse(response, 'failed to load move due to server error');
  return response.body;
}

// ORDERS
export async function LoadOrders(ordersId) {
  const client = await getClient();
  const response = await client.apis.orders.showOrders({
    ordersId,
  });
  checkResponse(response, 'failed to load orders due to server error');
  return response.body;
}

// SERVICE MEMBER
export async function LoadServiceMember(serviceMemberId) {
  const client = await getClient();
  const response = await client.apis.service_members.showServiceMember({
    serviceMemberId,
  });
  checkResponse(response, 'failed to load service member due to server error');
  return response.body;
}

export async function UpdateServiceMember(serviceMemberId, payload) {
  const client = await getClient();
  const response = await client.apis.service_members.patchServiceMember({
    serviceMemberId,
    patchServiceMemberPayload: payload,
  });
  checkResponse(
    response,
    'failed to update service member due to server error',
  );
  return response.body;
}

// BACKUP CONTACT
export async function LoadBackupContacts(serviceMemberId) {
  const client = await getClient();
  const response = await client.apis.backup_contacts.indexServiceMemberBackupContacts(
    {
      serviceMemberId,
    },
  );
  checkResponse(response, 'failed to load backup contacts due to server error');
  return response.body;
}

export async function UpdateBackupContact(backupContactId, payload) {
  const client = await getClient();
  const response = await client.apis.backup_contacts.updateServiceMemberBackupContact(
    {
      backupContactId,
      updateServiceMemberBackupContactPayload: payload,
    },
  );
  checkResponse(response, 'failed to load backup contacts due to server error');
  return response.body;
}

// PPM
export async function LoadPPMs(moveId) {
  const client = await getClient();
  const response = await client.apis.ppm.indexPersonallyProcuredMoves({
    moveId,
  });
  checkResponse(response, 'failed to load PPMs due to server error');
  return response.body;
}

// Move status
export async function ApproveBasics(moveId) {
  const client = await getClient();
  const response = await client.apis.office.approveMove({
    moveId,
  });
  checkResponse(response, 'failed to approve move due to server error');
  return response.body;
}

// Get PPM SIT Estimate
export async function LoadPPMSITEstimate(
  moveDate,
  originZip,
  destinationZip,
  weightEstimate,
) {
  const client = await getClient();
  const response = await client.apis.ppm.showPPMSitEstimate({
    planned_move_date: moveDate,
    origin_zip: originZip,
    destination_zip: destinationZip,
    weight_estimate: weightEstimate,
  });
  checkResponse(response, 'failed to get SIT estimate due to server error');
  console.log('response.body: ', response.body);
  return response.body;
}
