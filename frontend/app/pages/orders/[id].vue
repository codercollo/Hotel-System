<script setup lang="ts">
definePageMeta({ requiresAuth: true });

const route = useRoute();
const { order, loading, error, get, cancel } = useOrders();
const {
  initiate,
  pollUntilDone,
  payment,
  loading: payLoading,
  error: payError,
} = usePayments();

const cancelling = ref(false);
const cancelError = ref<string | null>(null);
const selectedProvider = ref("stripe");
const phoneNumber = ref("");
const paymentInitiated = ref(false);
const paymentDone = ref(false);
// checkout_code comes back on InitiateResponse, not on the Payment record.
// Store it here after initiation so the template can display the STK push code.
const checkoutCode = ref<string | null>(null);

useHead({ title: "Booking Detail" });
onMounted(() => get(route.params.id as string));

// ── Formatters ───────────────────────────────────────────────────────────────
const formatPrice = (p: number) =>
  (p / 100).toLocaleString("en-US", {
    style: "currency",
    currency: "USD",
    maximumFractionDigits: 0,
  });

const formatDate = (d: string) =>
  new Date(d).toLocaleDateString("en-US", {
    month: "long",
    day: "numeric",
    year: "numeric",
  });

const formatDateTime = (d: string) =>
  new Date(d).toLocaleString("en-US", {
    month: "short",
    day: "numeric",
    year: "numeric",
    hour: "2-digit",
    minute: "2-digit",
  });

// ── Derived booking data ──────────────────────────────────────────────────────
const roomName = computed(
  () =>
    (order.value?.metadata?.room_name as string) ||
    order.value?.items[0]?.name ||
    "Room Booking",
);
const checkIn = computed(
  () =>
    (order.value?.metadata?.check_in as string) ||
    order.value?.created_at ||
    "",
);
const checkOut = computed(
  () =>
    (order.value?.metadata?.check_out as string) ||
    order.value?.updated_at ||
    "",
);
const guestCount = computed(
  () => (order.value?.metadata?.guests as number) || 1,
);

const nights = computed(() => {
  if (!checkIn.value || !checkOut.value)
    return order.value?.items[0]?.quantity ?? 1;
  return Math.round(
    (new Date(checkOut.value).getTime() - new Date(checkIn.value).getTime()) /
      86_400_000,
  );
});

const subtotal = computed(
  () => order.value?.items.reduce((acc, i) => acc + i.subtotal, 0) ?? 0,
);
const taxes = computed(() =>
  Math.round(((order.value?.total ?? 0) * 0.15) / 1.15),
);

// ── Timeline ─────────────────────────────────────────────────────────────────
const timeline = computed(() => {
  if (!order.value) return [];
  const o = order.value;
  const paid = paymentDone.value && payment.value?.status === "completed";
  return [
    { label: "Booking Placed", date: formatDateTime(o.created_at), done: true },
    {
      label: "Payment Confirmed",
      date: paid ? formatDateTime(payment.value!.updated_at) : "—",
      done: paid,
    },
    {
      label: "Booking Confirmed",
      date: formatDateTime(o.updated_at),
      done: ["confirmed", "processing", "completed"].includes(o.status),
    },
    {
      label: "Check-In",
      date: checkIn.value ? formatDate(checkIn.value) : "—",
      done: o.status === "completed",
    },
    {
      label: "Check-Out",
      date: checkOut.value ? formatDate(checkOut.value) : "—",
      done: o.status === "completed",
    },
  ];
});

// ── Payment gating ───────────────────────────────────────────────────────────
// M-Pesa requires a phone number; other providers redirect or use saved cards.
const needsMpesaPhone = computed(() => selectedProvider.value === "mpesa");

const canPay = computed(
  () =>
    order.value?.status === "pending" &&
    !paymentInitiated.value &&
    (!needsMpesaPhone.value || phoneNumber.value.trim().length >= 10),
);

// ── Handlers ─────────────────────────────────────────────────────────────────
async function handlePay() {
  if (!order.value) return;
  try {
    const res = await initiate(
      order.value.id,
      selectedProvider.value,
      needsMpesaPhone.value ? phoneNumber.value.trim() : undefined,
    );
    paymentInitiated.value = true;
    checkoutCode.value = res.checkout_code ?? null;
    if (res.checkout_url) {
      window.open(res.checkout_url, "_blank");
    }

    // STK-push providers (M-Pesa) show the checkout code and then poll
    // All providers poll until the webhook confirms the terminal status.
    pollUntilDone(res.payment_id, async (p) => {
      paymentDone.value = true;
      // Refresh the order so order.status reflects the webhook update
      await get(order.value!.id);
    });
  } catch {
    // error is already set inside the composable
  }
}

const handleCancel = async () => {
  if (!order.value) return;
  cancelling.value = true;
  cancelError.value = null;
  try {
    await cancel(order.value.id);
  } catch (e: any) {
    cancelError.value = e.message;
  } finally {
    cancelling.value = false;
  }
};
</script>

<template>
  <div class="max-w-4xl mx-auto section-px py-12">
    <nav class="flex items-center gap-2 text-xs font-sans text-muted mb-8">
      <NuxtLink to="/orders" class="hover:text-forest transition-colors"
        >My Bookings</NuxtLink
      >
      <Icon name="lucide:chevron-right" class="w-3 h-3" />
      <span class="text-brand-900">{{ route.params.id }}</span>
    </nav>

    <!-- Loading -->
    <div v-if="loading" class="grid grid-cols-1 lg:grid-cols-3 gap-8">
      <div class="lg:col-span-2 space-y-6">
        <UiSkeleton class="h-48 rounded-2xl" />
        <UiSkeleton class="h-64 rounded-2xl" />
      </div>
      <div class="space-y-4">
        <UiSkeleton class="h-40 rounded-2xl" />
        <UiSkeleton class="h-32 rounded-2xl" />
      </div>
    </div>

    <UiAlert v-else-if="error" variant="error" :message="error" />

    <div v-else-if="order" class="grid grid-cols-1 lg:grid-cols-3 gap-8">
      <!-- ── Left column ─────────────────────────────────────────────────── -->
      <div class="lg:col-span-2 space-y-6">
        <!-- Header card -->
        <div class="card p-6">
          <div class="flex items-start justify-between mb-4">
            <div>
              <OrderStatusBadge :status="order.status" class="mb-3" />
              <h1
                class="font-display text-display-sm text-brand-900 font-light"
              >
                {{ roomName }}
              </h1>
              <p class="text-xs font-mono text-muted mt-1">{{ order.id }}</p>
            </div>
            <div class="text-right">
              <div class="font-display text-2xl text-brand-900">
                {{ formatPrice(order.total) }}
              </div>
              <div class="text-xs text-muted font-sans">total charged</div>
            </div>
          </div>

          <div
            class="grid grid-cols-2 sm:grid-cols-4 gap-4 pt-4 border-t border-surface-200"
          >
            <div>
              <p
                class="text-xs font-sans uppercase tracking-widest text-muted mb-1"
              >
                Check In
              </p>
              <p class="text-sm font-sans text-brand-900 font-medium">
                {{ checkIn ? formatDate(checkIn) : "—" }}
              </p>
            </div>
            <div>
              <p
                class="text-xs font-sans uppercase tracking-widest text-muted mb-1"
              >
                Check Out
              </p>
              <p class="text-sm font-sans text-brand-900 font-medium">
                {{ checkOut ? formatDate(checkOut) : "—" }}
              </p>
            </div>
            <div>
              <p
                class="text-xs font-sans uppercase tracking-widest text-muted mb-1"
              >
                Nights
              </p>
              <p class="text-sm font-sans text-brand-900 font-medium">
                {{ nights }}
              </p>
            </div>
            <div>
              <p
                class="text-xs font-sans uppercase tracking-widest text-muted mb-1"
              >
                Guests
              </p>
              <p class="text-sm font-sans text-brand-900 font-medium">
                {{ guestCount }}
              </p>
            </div>
          </div>
        </div>

        <!-- Timeline -->
        <div class="card p-6">
          <h3 class="font-display text-xl text-brand-900 mb-6">
            Booking Timeline
          </h3>
          <OrderTimeline :steps="timeline" />
        </div>

        <!-- Payment status card (shown after initiating) -->
        <div v-if="paymentInitiated && payment" class="card p-6">
          <h3 class="font-display text-xl text-brand-900 mb-4">Payment</h3>
          <PaymentStatusCard
            :status="payment.status"
            :amount="payment.amount"
            :provider="payment.provider"
            :ref="payment.provider_ref"
          />
          <!-- M-Pesa: show STK push code while waiting -->
          <p
            v-if="checkoutCode && !paymentDone"
            class="text-xs font-mono text-muted mt-2"
          >
            STK Push Code: {{ checkoutCode }}
          </p>
          <p
            v-if="!paymentDone"
            class="text-xs text-muted font-sans mt-3 flex items-center gap-1.5"
          >
            <Icon name="lucide:loader-2" class="w-3 h-3 animate-spin" />
            Waiting for payment confirmation…
          </p>
        </div>
      </div>

      <!-- ── Right column ────────────────────────────────────────────────── -->
      <div class="space-y-4">
        <OrderSummary
          :room="roomName"
          :check-in="checkIn"
          :check-out="checkOut"
          :nights="nights"
          :subtotal="subtotal"
          :taxes="taxes"
          :total="order.total"
        />

        <!-- Pay now panel — only for pending, unpaid orders -->
        <div v-if="order.status === 'pending' && !paymentDone" class="card p-5">
          <h3 class="font-display text-lg text-brand-900 mb-4">
            Complete Payment
          </h3>

          <PaymentMethodSelector v-model="selectedProvider" class="mb-4" />

          <!-- M-Pesa phone input -->
          <div v-if="needsMpesaPhone" class="mb-4">
            <label
              class="block text-xs uppercase tracking-widest font-sans text-muted mb-1.5"
            >
              M-Pesa Phone Number
            </label>
            <input
              v-model="phoneNumber"
              type="tel"
              placeholder="2547XXXXXXXX"
              class="input"
            />
          </div>

          <UiAlert
            v-if="payError"
            variant="error"
            :message="payError"
            class="mb-3 text-xs"
          />

          <UiButton
            class="w-full justify-center"
            :disabled="!canPay || payLoading"
            :loading="payLoading || (paymentInitiated && !paymentDone)"
            @click="handlePay"
          >
            <Icon name="lucide:lock" class="w-3.5 h-3.5" />
            {{
              paymentInitiated
                ? "Processing…"
                : `Pay ${formatPrice(order.total)}`
            }}
          </UiButton>

          <p class="text-[10px] text-muted font-sans text-center mt-2">
            Secured · SSL encrypted
          </p>
        </div>

        <!-- Payment success state -->
        <div
          v-if="paymentDone && payment?.status === 'completed'"
          class="card p-5 text-center"
        >
          <Icon
            name="lucide:check-circle"
            class="w-10 h-10 text-green-600 mx-auto mb-2"
          />
          <p class="font-display text-lg text-brand-900">Payment Complete</p>
          <p class="text-xs text-muted font-sans mt-1">
            Your booking is confirmed.
          </p>
        </div>

        <!-- Actions -->
        <div class="card p-5 space-y-3">
          <h3 class="font-display text-lg text-brand-900 mb-1">Actions</h3>
          <UiAlert
            v-if="cancelError"
            variant="error"
            :message="cancelError"
            class="text-xs"
          />
          <UiButton variant="outline" size="sm" class="w-full justify-center">
            <Icon name="lucide:download" class="w-3.5 h-3.5" /> Download Receipt
          </UiButton>
          <UiButton
            v-if="order.status === 'confirmed' || order.status === 'pending'"
            variant="outline"
            size="sm"
            :loading="cancelling"
            class="w-full justify-center text-red-600 border-red-200 hover:bg-red-50 hover:border-red-300"
            @click="handleCancel"
          >
            <Icon name="lucide:x-circle" class="w-3.5 h-3.5" /> Cancel Booking
          </UiButton>
        </div>

        <div class="card p-5">
          <h3 class="font-display text-lg text-brand-900 mb-3">Need Help?</h3>
          <p class="text-xs text-muted font-sans mb-4 leading-relaxed">
            Our concierge team is available 24/7 for any questions about your
            stay.
          </p>
          <a href="tel:+10000000000">
            <UiButton size="sm" class="w-full justify-center">
              <Icon name="lucide:phone" class="w-3.5 h-3.5" /> Call Concierge
            </UiButton>
          </a>
        </div>
      </div>
    </div>
  </div>
</template>
