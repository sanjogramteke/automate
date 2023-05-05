#shellcheck disable=SC2034
#shellcheck disable=SC2039
#shellcheck disable=SC2154
#stable channel

pkg_name="automate-cs-ocid"
pkg_description="Wrapper package for chef/ocid"
pkg_origin="chef"
# WARNING: Version managed by .expeditor/update_chef_server.sh
pkg_version="15.4.0"
vendor_origin="chef"
pkg_maintainer="Chef Software Inc. <support@chef.io>"
pkg_license=("Chef-MLSA")
pkg_upstream_url="https://www.chef.io/automate"
pkg_deps=(
  chef/mlsa
  "${local_platform_tools_origin:-chef}/automate-platform-tools"
  # WARNING: Version pin managed by .expeditor/update_chef_server.sh
  "${vendor_origin}/oc_id/15.4.0/20230105061030"
)

pkg_binds=(
  [automate-pg-gateway]="port"
  [pg-sidecar-service]="port"
)
pkg_exports=(
  [http-host]=network.host
  [http-port]=network.port
)

pkg_exposes=(http-port)

pkg_scaffolding="${local_scaffolding_origin:-chef}/automate-scaffolding"
automate_scaffolding_include_templates=(sqerl.config)

do_prepare() {
  GO_LDFLAGS="-X main.RubyPath=$(hab pkg path 'core/ruby27')"
  export GO_LDFLAGS

  build_line "Setting link for /usr/bin/env to 'coreutils'"
  [[ ! -f /usr/bin/env ]] && ln -s "$(pkg_path_for coreutils)/bin/env" /usr/bin/env
  return 0
}

do_download() {
  return 0
}

do_build() {
  return 0
}

do_install() {
  cd "$(hab pkg path 'chef/oc_id')/oc_id"

  export DATABASE_URL="postgresql://automate@127.0.0.1:5432/automate-cs-ocid?sslmode=verify-ca&sslcert=/hab/svc/automate-postgresql/config/server.crt&sslkey=/hab/svc/automate-postgresql/config/server.key&sslrootcert=/hab/svc/automate-postgresql/config/root.crt"

  # TODO :: Remove following line once new ocid hab package is used
  echo "gem 'tzinfo-data'" >> Gemfile

  export BUNDLE_SILENCE_ROOT_WARNING=1 GEM_PATH
  build_line "Setting BUNDLE_SILENCE_ROOT_WARNING=$BUNDLE_SILENCE_ROOT_WARNING"

  bundle config path "vendor/bundle"

  # TODO: Check if it works with path flag
  bundle install

  bundle exec bin/rake db:create
  bundle exec bin/rake db:migrate

  mkdir -p tmp
  chmod 777 -R tmp

  return 0
}
