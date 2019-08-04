##
## VARIABLES
##
# Program version.
__TAG = $(shell git describe --tags 2> /dev/null)
ifneq ($(__TAG),)
	VERSION ?= $(shell echo "$(__TAG)" | cut -c 2-)
else
	VERSION ?= undefined
endif


##
## TARGETS
##
# Generic:
.PHONY: default version setup install build clean run lint test review help
__ARGS = $(filter-out $@,$(MAKECMDGOALS))

default: run
version: # Show project version (derived from 'git describe').
	@echo $(VERSION)

setup: # Set this project up on a new environment.
	@$(MAKE) js-setup -- $(__ARGS)
install: # Install project dependencies.
	@$(MAKE) js-install -- $(__ARGS)

run: # Run project (development).
	@$(MAKE) js-run -- $(__ARGS)
build: # Build project.
	@$(MAKE) vue-build -- $(__ARGS)
clean: # Clean build artifacts.
	@$(MAKE) js-clean && \
	 $(MAKE) vue-clean

lint: # Lint and check code.
	@$(MAKE) js-lint -- $(__ARGS)
test: # Run tests.
	@echo "No testing procedure defined."
review: # Lint code and run tests.
	@$(MAKE) js-review -- $(__ARGS)

# Show usage for the targets in this Makefile.
help:
	@grep -E '^[a-zA-Z_-]+:.*?# .*$$' $(MAKEFILE_LIST) | \
	 awk 'BEGIN {FS = ":.*?# "}; {printf "\033[36m%-15s\033[0m %s\n", $$1, $$2}'


# Javascript:
.PHONY: js-setup js-install js-build js-run js-test js-lint js-review js-clean \
        js-release

js-setup:
	@$(MAKE) js-install -- $(__ARGS)
js-install:
	@echo "Installing package dependencies..." && yarn install $(__ARGS)
js-build:
	@echo "Building static files for production..." && yarn build $(__ARGS)
js-run:
	@echo "Starting development server..." && yarn dev $(__ARGS)
js-lint:
	@echo "Linting code with 'yarn lint'..." && yarn lint $(__ARGS)
js-review:
	@$(MAKE) js-lint -- $(__ARGS)
js-clean:
	@echo "Cleaning 'node_modules'..." && rm -rf node_modules && echo done


# Vue:
.PHONY: vue-build vue-clean
vue-build: # Compile Vue app for production.
	@echo "Build Vue app for production..." && yarn build:prod && echo done

vue-clean: # Clean up dist folder.
	@echo "Cleaning .dist/..." && rm -rf .dist/ && echo done


# HACKS:
# These targets are hacks that allow for Make targets to receive
# pseudo-arguments.
.PHONY: __FORCE
__FORCE:
%: __FORCE
	@:
