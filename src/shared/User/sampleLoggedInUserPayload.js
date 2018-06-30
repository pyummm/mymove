export const emptyPayload = {
  type: 'GET_LOGGED_IN_USER_SUCCESS',
  payload: {
    created_at: '2018-05-20T14:38:57.353Z',
    id: 'b46e651e-9d1c-4be5-bb88-bba58e817696',
    updated_at: '2018-05-20T14:38:57.353Z',
  },
};
export default {
  type: 'GET_LOGGED_IN_USER_SUCCESS',
  payload: {
    created_at: '2018-05-20T14:38:57.353Z',
    id: 'b46e651e-9d1c-4be5-bb88-bba58e817696',
    service_member: {
      affiliation: 'ARMY',
      backup_contacts: [
        {
          createdAt: '2018-05-31T00:02:57.302Z',
          email: 'foo@bar.com',
          id: '03b2979d-8046-437b-a6e4-11dbe251a912',
          name: 'Foo',
          permission: 'NONE',
          updated_at: '2018-05-31T00:02:57.302Z',
        },
      ],
      backup_mailing_address: {
        city: 'Washington',
        postal_code: '20021',
        state: 'DC',
        street_address_1: '200 K St',
      },
      created_at: '2018-05-25T15:48:49.918Z',
      current_station: {
        address: {
          city: 'Colorado Springs',
          country: 'United States',
          postal_code: '80913',
          state: 'CO',
          street_address_1: 'n/a',
        },
        affiliation: 'ARMY',
        created_at: '2018-05-20T18:36:45.034Z',
        id: '28f63a9d-8fff-4a0f-84ef-661c5c8c354e',
        name: 'Ft Carson',
        updated_at: '2018-05-20T18:36:45.034Z',
      },
      edipi: '1234567890',
      email_is_preferred: false,
      first_name: 'Erin',
      has_social_security_number: true,
      id: '1694e00e-17ff-43fe-af6d-ab0519a18ff2',
      is_profile_complete: true,
      last_name: 'Stanfill',
      middle_name: '',
      orders: [
        {
          created_at: '2018-05-27T21:36:10.219Z',
          has_dependents: false,
          id: '51953e97-25a7-430c-ba6d-3bd980a38b00',
          issue_date: '2018-05-11',
          status: 'CANCELED',
          moves: [
            {
              created_at: '2018-05-27T21:36:10.235Z',
              id: '593cc830-1a3e-44b3-ba5a-8809f02d000',
              locator: 'BLABLA',
              orders_id: '51953e97-25a7-430c-ba6d-3bd980a38b00',
              personally_procured_moves: [
                {
                  destination_postal_code: '76127',
                  incentive_estimate_min: 1495409,
                  incentive_estimate_max: 1652821,
                  has_additional_postal_code: false,
                  has_requested_advance: false,
                  has_sit: false,
                  id: 'cd67c9e4-ef59-45e5-94bc-767aaafe559e',
                  pickup_postal_code: '80913',
                  planned_move_date: '2018-06-28',
                  size: 'L',
                  status: 'CANCELED',
                  weight_estimate: 9000,
                },
              ],
              selected_move_type: 'PPM',
              status: 'CANCELED',
            },
          ],
          new_duty_station: {
            address: {
              city: 'Fort Worth',
              country: 'United States',
              postal_code: '76127',
              state: 'TX',
              street_address_1: 'n/a',
            },
            affiliation: 'NAVY',
            created_at: '2018-05-20T18:36:45.034Z',
            id: '44db8bfb-db7c-4c8d-bc08-5d683c4469ed',
            name: 'NAS Fort Worth',
            updated_at: '2018-05-20T18:36:45.034Z',
          },
          orders_type: 'PERMANENT_CHANGE_OF_STATION',
          report_by_date: '2018-05-29',
          service_member_id: '1694e00e-17ff-43fe-af6d-ab0519a18ff2',
          updated_at: '2018-05-25T21:39:02.429Z',
          uploaded_orders: {
            id: '24f18674-eec7-4c1f-b8c0-cb343a8c4f77',
            name: 'uploaded_orders',
            service_member_id: '1694e00e-17ff-43fe-af6d-ab0519a18ff2',
            uploads: [
              {
                bytes: 3932969,
                content_type: 'image/jpeg',
                created_at: '2018-05-25T21:38:06.235Z',
                filename: 'last vacccination.jpg',
                id: 'd56df2e3-1481-4dff-9a02-ef5c6bcae491',
                updated_at: '2018-05-25T21:38:06.235Z',
                url:
                  '/storage/documents/24f18674-eec7-4c1f-b8c0-cb343a8c4f77/uploads/d56df2e3-1481-4dff-9a02-ef5c6bcae491?contentType=image%2Fjpeg',
              },
              {
                bytes: 58036,
                content_type: 'image/png',
                created_at: '2018-05-25T21:38:57.655Z',
                filename: 'image (2).png',
                id: 'e2010a83-ac1e-45a2-9eb1-4e144c443c41',
                updated_at: '2018-05-25T21:38:57.655Z',
                url:
                  '/storage/documents/24f18674-eec7-4c1f-b8c0-cb343a8c4f77/uploads/e2010a83-ac1e-45a2-9eb1-4e144c443c41?contentType=image%2Fpng',
              },
            ],
          },
        },
        {
          created_at: '2018-05-25T21:36:10.219Z',
          has_dependents: false,
          id: '51953e97-25a7-430c-ba6d-3bd980a38b71',
          issue_date: '2018-05-11',
          status: 'DRAFT',
          moves: [
            {
              created_at: '2018-05-25T21:36:10.235Z',
              id: '593cc830-1a3e-44b3-ba5a-8809f02dfa7d',
              locator: 'WUMGLQ',
              orders_id: '51953e97-25a7-430c-ba6d-3bd980a38b71',
              personally_procured_moves: [
                {
                  destination_postal_code: '76127',
                  incentive_estimate_min: 1495409,
                  incentive_estimate_max: 1652821,
                  has_additional_postal_code: false,
                  has_requested_advance: false,
                  has_sit: false,
                  id: 'cd67c9e4-ef59-45e5-94bc-767aaafe559e',
                  pickup_postal_code: '80913',
                  planned_move_date: '2018-06-28',
                  size: 'L',
                  status: 'DRAFT',
                  weight_estimate: 9000,
                },
              ],
              selected_move_type: 'PPM',
              status: 'DRAFT',
            },
          ],
          new_duty_station: {
            address: {
              city: 'Fort Worth',
              country: 'United States',
              postal_code: '76127',
              state: 'TX',
              street_address_1: 'n/a',
            },
            affiliation: 'NAVY',
            created_at: '2018-05-20T18:36:45.034Z',
            id: '44db8bfb-db7c-4c8d-bc08-5d683c4469ed',
            name: 'NAS Fort Worth',
            updated_at: '2018-05-20T18:36:45.034Z',
          },
          orders_type: 'PERMANENT_CHANGE_OF_STATION',
          report_by_date: '2018-05-29',
          service_member_id: '1694e00e-17ff-43fe-af6d-ab0519a18ff2',
          updated_at: '2018-05-25T21:39:02.429Z',
          uploaded_orders: {
            id: '24f18674-eec7-4c1f-b8c0-cb343a8c4f77',
            name: 'uploaded_orders',
            service_member_id: '1694e00e-17ff-43fe-af6d-ab0519a18ff2',
            uploads: [
              {
                bytes: 3932969,
                content_type: 'image/jpeg',
                created_at: '2018-05-25T21:38:06.235Z',
                filename: 'last vacccination.jpg',
                id: 'd56df2e3-1481-4dff-9a02-ef5c6bcae491',
                updated_at: '2018-05-25T21:38:06.235Z',
                url:
                  '/storage/documents/24f18674-eec7-4c1f-b8c0-cb343a8c4f77/uploads/d56df2e3-1481-4dff-9a02-ef5c6bcae491?contentType=image%2Fjpeg',
              },
              {
                bytes: 58036,
                content_type: 'image/png',
                created_at: '2018-05-25T21:38:57.655Z',
                filename: 'image (2).png',
                id: 'e2010a83-ac1e-45a2-9eb1-4e144c443c41',
                updated_at: '2018-05-25T21:38:57.655Z',
                url:
                  '/storage/documents/24f18674-eec7-4c1f-b8c0-cb343a8c4f77/uploads/e2010a83-ac1e-45a2-9eb1-4e144c443c41?contentType=image%2Fpng',
              },
            ],
          },
        },
      ],
      personal_email: 'erin@truss.works',
      phone_is_preferred: true,
      rank: 'O_4_W_4',
      residential_address: {
        city: 'Somewhere',
        postal_code: '80913',
        state: 'CO',
        street_address_1: '123 Main',
      },
      telephone: '555-555-5556',
      text_message_is_preferred: true,
      updated_at: '2018-05-25T21:39:10.484Z',
      user_id: 'b46e651e-9d1c-4be5-bb88-bba58e817696',
    },
    updated_at: '2018-05-20T14:38:57.353Z',
  },
};
