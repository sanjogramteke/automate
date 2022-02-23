describe('chef server details', () => {
  let adminIdToken = '';
  const now = Cypress.moment().format('MMDDYYhhmm');
  const cypressPrefix = 'infra';
  const serverName = `${cypressPrefix} server ${now}`;
  const updatedServerName = `${cypressPrefix} updated server ${now}`;
  const serverID = serverName.split(' ').join('-');
  const serverFQDN = Cypress.env('AUTOMATE_INFRA_SERVER_FQDN');
  const webuiKey = Cypress.env('AUTOMATE_INFRA_WEBUI_KEY').replace(/\\n/g, '\n');
  const orgName = `${cypressPrefix} org ${now}`;
  const generatedOrgID = orgName.split(' ').join('-');
  const customOrgID = `${cypressPrefix}-custom-id-${now}`;
  const adminUser = 'test_admin_user';

  // using dummy admin key value for creating the org
  const adminKey = 'Dummy--admin--key';

  before(() => {
    cy.adminLogin('/').then(() => {
      const admin = JSON.parse(<string>localStorage.getItem('chef-automate-user'));
      adminIdToken = admin.id_token;

      cy.request({
        auth: { bearer: adminIdToken },
        method: 'POST',
        url: '/api/v0/infra/servers',
        body: {
          id: serverID,
          name: serverName,
          fqdn: serverFQDN,
          ip_address: '',
          webui_key: webuiKey
        }
      });

      cy.visit(`/infrastructure/chef-servers/${serverID}`);
      cy.get('app-welcome-modal').invoke('hide');
    });

    cy.restoreStorage();
  });

  beforeEach(() => {
    cy.restoreStorage();
  });

  afterEach(() => {
    cy.saveStorage();
  });

  describe('chef server details page', () => {
    it('displays server details', () => {
      cy.get('chef-breadcrumbs').contains('Chef Infra Servers');
      cy.get('chef-breadcrumbs').contains(serverName);
      cy.get('.page-title').contains(serverName);
      cy.get('[data-cy=add-org-button]').contains('Add Chef Organization');
    });

    it('it can update the webui key', () => {
      cy.get('[data-cy=update-web-ui-key]').contains('Update').click();
      cy.get('app-chef-server-details .sidenav-header').should('exist');
      cy.get('[data-cy=enter-webui-key]').type(webuiKey);

      cy.get('[data-cy=update-webui-key-button]').click();
      cy.get('app-chef-server-details .sidenav-header').should('not.be.visible');
    });

    it('can validate the server webui key', () => {
      cy.get('[data-cy=valid-webui-key]').contains('Valid');
    });

    it('can check empty org list', () => {
      cy.get('#org-table-container chef-th').should('not.be.visible');
      cy.get('[data-cy=empty-list]').should('be.visible');
      cy.get('[data-cy=no-records]').contains('No Organization available');
    });

    it('can add a organization', () => {
      cy.get('[data-cy=add-org-button]').contains('Add Chef Organization')
        .click({multiple: true, force: true});
      cy.get('app-chef-server-details chef-modal').should('exist');
      cy.get('[data-cy=org-name]').type(orgName);
      cy.get('[data-cy=id-label]').contains(generatedOrgID);
      cy.get('[data-cy=admin-user]').type(adminUser);
      cy.get('[data-cy=admin-key]').clear().invoke('val', adminKey).trigger('input');

      cy.get('[data-cy=add-button]').click();
      cy.get('app-chef-server-details chef-modal').should('not.be.visible');

      // verify success notification and then dismiss it
      // so it doesn't get in the way of subsequent interactions
      cy.get('app-notification.info').should('be.visible');
      cy.get('app-notification.info chef-icon').click({multiple: true, force: true});
      cy.get('app-chef-server-details chef-tbody chef-td').contains(orgName).should('exist');
    });

    it('lists organizations', () => {
      cy.get('[data-cy=add-org-button]').contains('Add Chef Organization');
      cy.get('#org-table-container chef-th').contains('Name');
      cy.get('#org-table-container chef-th').contains('Admin');
      cy.get('#org-table-container chef-th').contains('Projects');
    });

    it('can create a chef organization with a custom ID', () => {
      cy.get('[data-cy=add-org-button]').contains('Add Chef Organization')
        .click({multiple: true, force: true});
      cy.get('app-chef-server-details chef-modal').should('exist');
      cy.get('[data-cy=org-name]').type(orgName);
      cy.get('[data-cy=add-id]').should('not.be.visible');
      cy.get('[data-cy=edit-button]').contains('Edit ID').click();
      cy.get('[data-cy=id-label]').should('not.be.visible');
      cy.get('[data-cy=add-id]').should('be.visible').clear().type(customOrgID);
      cy.get('[data-cy=admin-user]').type(adminUser);
      cy.get('[data-cy=admin-key]').clear().invoke('val', adminKey).trigger('input');

      cy.get('[data-cy=add-button]').click();
      cy.get('app-chef-server-details chef-modal').should('not.be.visible');

      // verify success notification and then dismiss it
      // so it doesn't get in the way of subsequent interactions
      cy.get('app-notification.info').should('be.visible');
      cy.get('app-notification.info chef-icon').click({multiple: true, force: true});

      cy.get('app-chef-server-details chef-tbody chef-td').contains(orgName).should('exist');
    });

    it('fails to create a chef organization with a duplicate ID', () => {
      cy.get('[data-cy=add-org-button]').contains('Add Chef Organization').click();
      cy.get('app-chef-server-details chef-modal').should('exist');
      cy.get('[data-cy=org-name]').type(orgName);
      cy.get('[data-cy=add-id]').should('not.be.visible');
      cy.get('[data-cy=edit-button]').contains('Edit ID').click();
      cy.get('[data-cy=id-label]').should('not.be.visible');
      cy.get('[data-cy=add-id]').should('be.visible').clear().type(customOrgID);
      cy.get('[data-cy=admin-user]').type(adminUser);
      cy.get('[data-cy=admin-key]').clear().invoke('val', adminKey).trigger('input');
      cy.get('[data-cy=add-button]').click();
      cy.get('app-chef-server-details chef-modal chef-error').contains('already exists')
        .should('be.visible');

      //  here we exit with the chef-modal exit button in the top right corner
      cy.get('app-chef-server-details chef-modal chef-button.close').first().click();
    });

    it('can cancel creating a chef organization', () => {
      cy.get('[data-cy=add-org-button]').contains('Add Chef Organization').click();
      cy.get('app-chef-server-details chef-modal').should('exist');

      // scoll to show the cancel button
      cy.get('[data-cy=cancel-button]').scrollIntoView();

      // here we exit with the Cancel button
      cy.get('[data-cy=cancel-button]').contains('Cancel').should('be.visible').click();
      cy.get('app-chef-server-details  chef-modal').should('not.be.visible');
    });

    it('can delete a chef organization', () => {
      cy.get('app-chef-server-details chef-td a').contains(orgName).parent().parent()
        .find('.mat-select-trigger').as('controlMenu');

      // we throw in a `should` so cypress retries until introspection allows menu to be shown
      cy.get('@controlMenu').scrollIntoView().should('be.visible').click();
      cy.get('[data-cy=delete-org]').should('be.visible').click();

      // accept dialog
      cy.get('app-chef-server-details chef-button').contains('Delete').click();

      // verify success notification and then dismiss it
      cy.get('app-notification.info').should('be.visible');
      cy.get('app-notification.info chef-icon').click({multiple: true, force: true});

      cy.get('app-chef-server-details chef-tbody chef-td')
        .contains(customOrgID).should('not.exist');
    });

    it('can check create organization button is disabled until all inputs are filled in', () => {
      cy.get('[data-cy=add-org-button]').contains('Add Chef Organization').click();
      cy.get('app-chef-server-details chef-modal').should('exist');
      cy.get('[data-cy=org-name]').type(orgName);
      cy.get('[data-cy=id-label]').contains(generatedOrgID);
      cy.get('[data-cy=admin-key]').clear().invoke('val', adminKey).trigger('input');

      // check for disabled
      cy.get('[data-cy=add-button]')
      .invoke('attr', 'disabled')
      .then(disabled => {
        disabled ? cy.log('buttonIsDiabled') : cy.get('[data-cy=add-button]').click();
      });

      cy.get('app-chef-server-details chef-modal').should('exist');

      // here we exit with the Cancel button
      cy.get('chef-button').contains('Cancel').should('be.visible').click();
      cy.get('app-chef-server-details  chef-modal').should('not.be.visible');
    });

    // details tabs specs
    it('can can switch to details tab', () => {
      cy.get('[data-cy=details-tab]').contains('Details').click();
    });

    it('can check save button is disabled until server details input value is not changed', () => {
      cy.get('[data-cy=update-server-name]').clear().type(serverName);
      cy.get('[data-cy=update-server-fqdn]').clear().type(serverFQDN);

      // check for disabled save button
      cy.get('[data-cy=save-button]')
      .invoke('attr', 'disabled')
      .then(disabled => {
        disabled ? cy.log('buttonIsDiabled') : cy.get('[data-cy=save-button]').click();
      });
    });

    xit('can check save button is disabled until all inputs are filled in', () => {
      cy.get('[data-cy=update-server-name]').clear().type(serverName);
      cy.get('[data-cy=update-server-fqdn]').clear().type(serverFQDN);

      // check for disabled save button
      cy.get('[data-cy=save-button]')
      .invoke('attr', 'disabled')
      .then(disabled => {
        disabled ? cy.log('buttonIsDiabled') : cy.get('[data-cy=save-button]').click();
      });

      cy.get('app-chef-server-details chef-error').contains('is required')
        .should('be.visible');
    });

    it('can update the server', () => {
      cy.get('[data-cy=update-server-name]').clear().type(updatedServerName);
      cy.get('[data-cy=update-server-fqdn]').clear().type(serverFQDN);
      cy.get('[data-cy=save-button]').click();

      cy.get('app-notification.info').contains('Successfully updated server');
      cy.get('app-notification.info chef-icon').click({ multiple: true });
    });

    xit('can update the server', () => {
      cy.get('[data-cy=update-server-name]').clear().type(updatedServerName);
      cy.get('chef-select').contains('FQDN').click();
      cy.get('chef-select chef-option').contains('IP Address').click();
      cy.get('[data-cy=update-server-ip-address]').clear().type(serverIP);
      cy.get('[data-cy=save-button]').click();

      cy.get('app-notification.info').contains('Successfully updated server');
      cy.get('app-notification.info chef-icon').click({ multiple: true });
    });
  });

  function getPreviewData(migration_id: string) {
    return cy.fixture('infra-proxy/previewData.json');
  }

  function checkResponse(response: any) {
    if (response.body.staged_data) {
      console.log('-----------checkResponse----------', response.body);

    }
  }

  describe('if migration process started', () => {
    let isStarted = true;
    beforeEach(() => {
      cy.fixture('infra-proxy/migrationStatus.json').then(data => {
        data.migrationStatus = data;
        data.migration_type = data.migrationStatus.migration_type;
        const migration_status = data.migrationStatus.migration_status;
      });
    });

    it('sync in progress visible', () => {
      const element = '[data-cy=lastStatus]';
      cy.get('body').then((body) => {
        if (body.find(element).length > 0) {
          isStarted = false;
            console.log('into', isStarted);
        }
      });
    });

    it('preview button is not available if migration is not in create preview mode', () => {
      if (isStarted === true) {
        cy.get('[data-cy=sync-button]').contains('Sync In Progress');
        cy.get('[data-cy=show-preview]').contains('Click to Preview').should('not.exist');
      }
    });

    it('check preview is available', () => {
      if (isStarted === true) {
        cy.get('[data-cy=lastStatus]').then(() => {
            cy.get('[data-cy=show-preview]').contains('Click to Preview').click();
            getPreviewData('').then((response) => {
              checkResponse(response);
            });
            cy.get('app-chef-server-details .sidenav-header').should('exist');

            cy.get('[data-cy=confirm-migration-button]').click();
            cy.get('app-chef-server-details .sidenav-header').should('not.be.visible');

            cy.get('app-notification.info').contains('Confirm preview successful.');
            cy.get('app-notification.info chef-icon').click({ multiple: true });
        });
      }
    });
  });
});
